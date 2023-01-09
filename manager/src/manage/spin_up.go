package manage

import (
	"context"
	"fmt"
	"manager/src/models"
	"manager/src/networking"
	"net"
	"os"
	"time"

	fk_sdk "github.com/firecracker-microvm/firecracker-go-sdk"
	fk_models "github.com/firecracker-microvm/firecracker-go-sdk/client/models"
)

func setupConfig(uuid string) fk_sdk.Config {
	var config fk_sdk.Config
	config.LogPath = "/home/cgeorge/genez/unikernel-function-runtime/manager/src/manage/log.log"
	// Config rootfs drive
	var drive_id string = "rootfs"
	var drive_path string = "/home/cgeorge/genez/unikernel-function-runtime/images/uuid1.img"
	var is_root bool = true
	var is_read_only bool = true

	var fkDrive fk_models.Drive = fk_models.Drive{
		DriveID:      &drive_id,
		PathOnHost:   &drive_path,
		IsReadOnly:   &is_read_only,
		IsRootDevice: &is_root,
	}

	config.Drives = []fk_models.Drive{fkDrive}

	// Config machine
	var vcpu_count int64 = 2
	var mem_size int64 = 1024
	var smt bool = true

	var fkMachine fk_models.MachineConfiguration = fk_models.MachineConfiguration{
		VcpuCount:  &vcpu_count,
		MemSizeMib: &mem_size,
		Smt:        &smt,
	}

	config.MachineCfg = fkMachine

	// Config network
	var fkNetIf fk_sdk.NetworkInterface
	var mac_addr string = networking.GenerateMACAddr_S()

	tap_dev_name := "tap_" + uuid
	fmt.Println(tap_dev_name)
	_, tap_net, _ := net.ParseCIDR("172.16.1.1/30")
	fkNetIf.StaticConfiguration = &fk_sdk.StaticNetworkConfiguration{
		MacAddress:  mac_addr,
		HostDevName: tap_dev_name,
		IPConfiguration: &fk_sdk.IPConfiguration{
			IPAddr:  *tap_net,
			Gateway: net.ParseIP("172.16.1.1"),
			IfName:  "eth0",
		},
	}

	config.SocketPath = "/tmp/fk_" + uuid
	config.NetworkInterfaces = []fk_sdk.NetworkInterface{fkNetIf}

	// Config kernel
	config.KernelImagePath = os.Getenv("KERNEL_PATH") + "/osv/osv-loader.elf.x86_64"

	// Config kernel networking
	config.KernelArgs += "--ip=eth0,172.16.1.2,255.255.255.252 --defaultgw=172.16.1.1 --nameserver=172.16.1.1 --nopci --rootfs=rofs java.so -cp /java-httpserver/user_app.jar:/java-httpserver/gson-2.10.jar:/java-httpserver HttpServerApp"

	// Command line

	return config
}

func NewInstance(request models.RunImageRequest, uuid string) (models.ExecutionResult, error) {
	// Cleanup when exiting this func
	os.Remove("/tmp/fk_" + "test")
	defer os.Remove("/tmp/fk_" + uuid)
	ctx := context.Background()
	// Create config for instance
	var fkConfig fk_sdk.Config = setupConfig(uuid)
	// Create output file
	stdoutPath := "/tmp/stdout_" + uuid + ".log"
	stdout, err := os.OpenFile(stdoutPath, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		panic(fmt.Errorf("failed to create stdout file: %v", err))
	}

	// Create new machine
	opts := fk_sdk.VMCommandBuilder{}.
		WithSocketPath("/tmp/fk_" + uuid).
		WithBin(os.Getenv("FIRECRACKER_PATH")).
		WithStdout(stdout).
		Build(ctx)

	new_machine, err := fk_sdk.NewMachine(ctx, fkConfig, fk_sdk.WithProcessRunner(opts), func(m *fk_sdk.Machine) {
		m.Handlers.FcInit = m.Handlers.FcInit.Swap(fk_sdk.Handler{
			Name: "fcinit.SetupKernelArgs",
			Fn: func(ctx context.Context, m *fk_sdk.Machine) error {
				fmt.Println("Kernel_ARG:")
				fmt.Println(m.Cfg.KernelArgs)
				return nil
			},
		})
	})
	if err != nil {
		fmt.Println(err)
		return models.ExecutionResult{}, err
	}
	tap_intf, err := networking.ConfigureTap(uuid)

	// Start execution and measure time from boot command to output return
	t1 := time.Now()
	new_machine.Start(ctx)
	fmt.Println(tap_intf.Name)
	if err != nil {
		fmt.Println("NETWORK:")
		fmt.Println(err)
		return models.ExecutionResult{}, err
	}

	// Execute function request
	function_output := ExecuteFunction(request, *tap_intf)
	elapsed := time.Since(t1)

	return models.ExecutionResult{
		Output:   function_output,
		Duration: elapsed.String(),
	}, nil
}
