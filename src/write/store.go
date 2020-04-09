package write

import (
	"log"
	"os"
)

// StoreEnv writes the Environtment variables to a file (environment)
func StoreEnv() {
	f, err := os.Create("environment")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()
	f.WriteString("CLUSTERDOMAIN=" + os.Getenv("CLUSTERDOMAIN") + "\n")
	f.WriteString("NAMECLUSTER=" + os.Getenv("NAMECLUSTER") + "\n")
	f.WriteString("NAMEDOMAIN=" + os.Getenv("NAMEDOMAIN") + "\n")
	f.WriteString("NAMEVPC=" + os.Getenv("NAMEVPC") + "\n")
	f.WriteString("AWSREGION1=" + os.Getenv("awsRegion1") + "\n")
	f.WriteString("AWSREGION2=" + os.Getenv("awsRegion2") + "\n")
	f.WriteString("AWSREGION3=" + os.Getenv("awsRegion3") + "\n")
	f.WriteString("DIRBASE=" + os.Getenv("DIRBASE") + "\n")

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
	os.Setenv("CLUSTERDOMAIN", "")
	os.Setenv("NAMECLUSTER", "")
	os.Setenv("NAMEDOMAIN", "")
	os.Setenv("NAMEVPC", "")

	// Enabled Regions
	// Hard coded for now
	os.Setenv("awsRegion1", "us-gov-west-1")
	os.Setenv("awsRegion2", "us-gov-west-1")
	os.Setenv("awsRegion3", "us-gov-west-1")

	// Working Variables
	os.Setenv("DIRBASE", "/root/PlatformOne")
	os.Setenv("DIRARTIFACTS", "")
}
