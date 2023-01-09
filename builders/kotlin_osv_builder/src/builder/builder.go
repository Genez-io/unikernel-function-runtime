package builder

import (
	"fmt"
	"kotlin_osv_builder/src/models"
	"os"
	"os/exec"
	"time"

	"github.com/google/uuid"
)

func UploadUnikernelMetadata(unikernel models.Unikernel) bool {
	fmt.Println(unikernel)
	return true
}

func BuildOSvImage(req_data models.CreateUnikernelRequest) (models.CreateUnikernelResponse, error) {
	// Generate UUID used for temporary files and unikernel identifier
	uuid_str := uuid.New().String()

	// Start measurement for unikernel image creation
	t1 := time.Now()

	// Build runtime and nanos unikernel config
	unikernel, err := BuildJavaKtRuntime(uuid_str, req_data.Code)
	if err != nil {
		return models.CreateUnikernelResponse{}, fmt.Errorf("failed at runtime create step")
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
	o, err := exec.Command(os.Getenv("OPS_PATH"), ops_args...).Output()
	fmt.Println(string(o))
	if err != nil {
		return models.CreateUnikernelResponse{}, fmt.Errorf("failed at ops step")
	}

	unikernel.RootFsImg = os.Getenv("ROOTFS_PATH") + uuid_str
	fmt.Println(unikernel.RootFsImg)
	_, err = exec.Command("mv", os.Getenv("IMAGES_PATH")+unikernel.RootFsImg, "/images").Output()
	if err != nil {
		return models.CreateUnikernelResponse{}, fmt.Errorf("failed at copy step")
	}

	os.Chmod("/images"+unikernel.RootFsImg, 0777)

	UploadUnikernelMetadata(*unikernel)
	t2 := time.Now()
	diff := t2.Sub(t1)

	// Remove temp files on function exit
	defer os.Remove("/tmp/ufr_code_" + uuid_str + ".jar")
	defer os.Remove("/tmp/ufr_config_" + uuid_str + ".json")
	defer os.Remove("/tmp/ufr_code_" + uuid_str + ".kt")
	defer os.RemoveAll("/tmp/ufr_jvm_runtime_" + uuid_str)

	return models.CreateUnikernelResponse{
		UUID:         uuid_str,
		CreationTime: fmt.Sprint(diff),
	}, nil
}
