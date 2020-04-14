package write

import (
	"log"
	"os"
)

//var dirBase = os.Getenv("HOME") + "PlatformOne"

// StoreEnv writes the Environtment variables to a file (environment)
func StoreEnv() {
	f, err := os.Create("environment")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()
	f.WriteString("dirBase=" + os.Getenv("HOME") + "/PlatformOne" + "\n")
	f.WriteString("dirArtifacts=" + os.Getenv("HOME") + "/PlatformOne" + os.Getenv("dirArtifacts") + "\n")
	f.WriteString("VPCNAME=" + os.Getenv("VPCNAME") + "\n")
	f.WriteString("CLUSTERNAME=" + os.Getenv("CLUSTERNAME") + "\n")
	f.WriteString("NAMEDOMAIN=" + os.Getenv("NAMEDOMAIN") + "\n")
	f.WriteString("CLUSTERDOMAIN=" + os.Getenv("CLUSTERDOMAIN") + "\n")
	f.WriteString("AWSREGION1=" + os.Getenv("AWSREGION1") + "\n")
	f.WriteString("AWSREGION2=" + os.Getenv("AWSREGION2") + "\n")
	f.WriteString("AWSREGION3=" + os.Getenv("AWSREGION3") + "\n")

	if err != nil {
		log.Fatal(err)
		f.Close()
		return
	}

	err = f.Close()
}

// CreateEnv writes defaults to the environment variables
func CreateEnv() {
	// User Defined Variables
	os.Setenv("CLUSTERDOMAIN", "")
	os.Setenv("CLUSTERNAME", "")
	os.Setenv("NAMEDOMAIN", "")
	os.Setenv("VPCNAME", "")

	// Enabled Regions
	// Hard coded for now
	os.Setenv("AWSREGION1", "us-gov-west-1")
	os.Setenv("AWSREGION2", "us-gov-west-1")
	os.Setenv("AWSREGION3", "us-gov-west-1")

	// Working Variables
	os.Setenv("dirBase", os.Getenv("HOME")+"/"+"PlatformOne")
	os.Setenv("dirArtifacts", "")
}
