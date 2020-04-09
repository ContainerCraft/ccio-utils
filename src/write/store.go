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
	f.WriteString("nameVpc=" + os.Getenv("nameVpc") + "\n")
	f.WriteString("nameCluster=" + os.Getenv("nameCluster") + "\n")
	f.WriteString("nameDomain=" + os.Getenv("nameDomain") + "\n")
	f.WriteString("clusterDomain=" + os.Getenv("clusterDomain") + "\n")
	f.WriteString("AWSREGION1=" + os.Getenv("awsRegion1") + "\n")
	f.WriteString("AWSREGION2=" + os.Getenv("awsRegion2") + "\n")
	f.WriteString("AWSREGION3=" + os.Getenv("awsRegion3") + "\n")

	if err != nil {
		log.Fatal(err)
		f.Close()
		return
	}

	err = f.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
}

// CreateEnv writes defaults to the environment variables
func CreateEnv() {
	// User Defined Variables
	os.Setenv("clusterDomain", "")
	os.Setenv("nameCluster", "")
	os.Setenv("nameDomain", "")
	os.Setenv("nameVpc", "")

	// Enabled Regions
	// Hard coded for now
	os.Setenv("awsRegion1", "us-gov-west-1")
	os.Setenv("awsRegion2", "us-gov-west-1")
	os.Setenv("awsRegion3", "us-gov-west-1")

	// Working Variables
	os.Setenv("dirBase", os.Getenv("HOME") + "/" + "PlatformOne")
	os.Setenv("dirArtifacts", "")
}
