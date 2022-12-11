package builder

import (
	"encoding/json"
	"fmt"
	"java_builder/src/models"
	"os"
	"os/exec"
	"regexp"
)

// func generateConfig(lang string, code_file string) (models.NanosConfig, error) {
// 	var config models.NanosConfig
// 	switch lang {
// 	case "swift":
// 		config.Program = "./" + code_file
// 		config.Files = []string{code_file}
// 	case "kotlin":
// 		config.Program = os.Getenv("JAVA_PATH")
// 		config.Args = []string{"-jar", code_file}
// 		config.Files = []string{code_file, os.Getenv("LIBSTDCPP_PATH"), os.Getenv("LIBM_PATH"), os.Getenv("LIBGCC_PATH")}
// 	case "nodejs":
// 		config.Program = os.Getenv("NODE_PATH")
// 		config.Args = []string{code_file}
// 		config.Files = []string{code_file}

// 	}
// 	return config, nil
// }

func BuildJavaKtRuntime(uuid string, code string) (*models.Unikernel, error) {
	// Create temporary file for code
	code_bytes := []byte(code)

	code_file := "/tmp/ufr_code_" + uuid + ".kt"
	err := os.WriteFile(code_file, code_bytes, 0644)
	if err != nil {
		return nil, err
	}

	// // Remove code file on function exit
	// defer os.Remove(code_file)

	var unikernel models.Unikernel
	// Build with Kotlinc
	kotlinc_args := []string{
		code_file,
		"-include-runtime",
		"-d",
		"/tmp/ufr_code_" + uuid + ".jar",
	}

	fmt.Println(kotlinc_args)
	kotlink_output, err := exec.Command(os.Getenv("KOTLINC_PATH"), kotlinc_args...).Output()
	fmt.Println("kotlinc:" + string(kotlink_output))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Analyse dependencies using jdeps
	jdeps_args := []string{
		"-s",
		"/tmp/ufr_code_" + uuid + ".jar",
	}

	// Grab raw dependency output
	jdeps_output, err := exec.Command("jdeps", jdeps_args...).Output()
	fmt.Println("jdeps:\n" + string(jdeps_output))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Filter output and store in dependency array
	re, err := regexp.Compile("(?: -> )(.*)")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var jdeps_result []string
	for _, v := range re.FindAllStringSubmatch(string(jdeps_output), -1) {
		jdeps_result = append(jdeps_result, v[1])
	}

	fmt.Println(jdeps_result)
	// Build custom runtime using jlink
	jlink_args := []string{"--add-modules"}
	var modules string = ""
	for i, v := range jdeps_result {
		modules += v
		if i != len(jdeps_result)-1 {
			modules += ","
		}
	}

	jlink_args = append(jlink_args, modules)
	jlink_args = append(jlink_args, "--output")
	jlink_args = append(jlink_args, "/tmp/ufr_jvm_runtime_"+uuid)
	jlink_args = append(jlink_args, "--no-header-files")
	jlink_args = append(jlink_args, "--no-man-pages")
	jlink_args = append(jlink_args, "--strip-debug")
	jlink_args = append(jlink_args, "--vm=server")

	fmt.Println(jlink_args)
	jlink_output, err := exec.Command("jlink", jlink_args...).Output()
	fmt.Println("jlink:" + string(jlink_output))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Generate nanos unikernel config
	unikernel.Config.Program = "/tmp/ufr_jvm_runtime_" + uuid + "/bin/java"

	unikernel.Config.Args = []string{
		"-jar",
		"/tmp/ufr_code_" + uuid + ".jar",
	}

	unikernel.Config.Files = []string{
		"/tmp/ufr_code_" + uuid + ".jar",
		os.Getenv("LIBSTDCPP_PATH"),
		os.Getenv("LIBZ_PATH"),
		os.Getenv("LIBM_PATH"),
		os.Getenv("LIBGCC_PATH"),
		"/usr/lib/x86_64-linux-gnu/librt.so.1",
	}

	// unikernel.Config.MapDirs = map[string]string{"/tmp/ufr_jvm_runtime_" + uuid: "/runtime"}
	unikernel.Config.Dirs = []string{"/tmp/ufr_jvm_runtime_" + uuid}

	unikernel.UUID = uuid
	unikernel.CreatedBy = "TODO"

	// Create temporary file for config
	// Marshalling results in empty config fields being ommited
	config_bytes, err := json.Marshal(unikernel.Config)
	if err != nil {
		return nil, err
	}
	unikernel.ConfigPath = "/tmp/ufr_config_" + uuid + ".json"
	err = os.WriteFile(unikernel.ConfigPath, config_bytes, 0644)
	if err != nil {
		return nil, err
	}

	return &unikernel, nil
}
