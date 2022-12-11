package builder

import (
	"fmt"
	"java_builder/src/models"
	"os"
	"os/exec"
	"time"
)

func UploadUnikernelMetadata(unikernel models.Unikernel) bool {
	fmt.Println(unikernel)
	return true
}

func BuildNanosImage(req_data models.CreateUnikernelRequest) (models.CreateUnikernelResponse, error) {
	// Generate UUID used for temporary files and unikernel identifier
	uuid_str := req_data.UUID

	t1 := time.Now()

	// Build runtime and nanos unikernel config
	unikernel, err := BuildJavaKtRuntime(uuid_str, req_data.Code)
	if err != nil {
		return models.CreateUnikernelResponse{}, err
	}
	// Run OPS and build image
	ops_args := []string{
		"image",
		"create",
		"-c",
		unikernel.ConfigPath,
		"-i",
		uuid_str,
	}

	fmt.Println(ops_args)
	out, err := exec.Command(os.Getenv("OPS_PATH"), ops_args...).Output()
	fmt.Println(string(out))
	if err != nil {
		fmt.Println(err)
		return models.CreateUnikernelResponse{}, err
	}

	unikernel.KernelImg = os.Getenv("KERNEL_PATH")
	unikernel.RootFsImg = os.Getenv("ROOTFS_PATH") + uuid_str
	UploadUnikernelMetadata(*unikernel)
	t2 := time.Now()
	diff := t2.Sub(t1)

	// Remove temp files on function exit
	defer os.Remove("/tmp/ufr_code_" + uuid_str + ".jar")
	defer os.Remove("/tmp/ufr_config_" + uuid_str + ".jar")
	defer os.Remove("/tmp/ufr_code_" + uuid_str + ".kt")
	defer os.RemoveAll("/tmp/ufr_jvm_runtime_" + uuid_str)

	return models.CreateUnikernelResponse{
		UUID:         uuid_str,
		CreationTime: fmt.Sprint(diff),
	}, nil
}
