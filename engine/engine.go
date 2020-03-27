package engine

import (
	_"fmt"
	"github.com/scott-x/gutils/cmd"
	"github.com/scott-x/gutils/fs"
	"github.com/scott-x/gutils/cl"
	"os/exec"
	"strings"
)

func Run(){
	home := fs.GetEnv("HOME");
	desktop := home + "/Desktop";
	target_folder := desktop + "/icons.iconset";
	cmd.AddQuestion("png_path", "please drag your image here:", "expect .png format: ", ".*.png")
	answers := cmd.Exec()
	full := strings.Trim(answers["png_path"]," ")

	fs.CreateDirIfNotExist(target_folder);
	  
	//command
	err :=exec.Command("sips","-z","16","16", full, "--out", target_folder+"/icon_16x16.png").Run()
	if err!=nil{
		panic(err)
	}
	err =exec.Command("sips","-z","32","32", full, "--out", target_folder+"/icon_16x16@2x.png").Run()
	if err!=nil{
		panic(err)
	}
	err =exec.Command("sips","-z","32","32", full, "--out", target_folder+"/icon_32x32.png").Run()
	if err!=nil{
		panic(err)
	}
	err =exec.Command("sips","-z","64","64", full, "--out", target_folder+"/icon_32x32@2x.png").Run()
	if err!=nil{
		panic(err)
	}
	err =exec.Command("sips","-z","64","64", full, "--out", target_folder+"/icon_64x64.png").Run()
	if err!=nil{
		panic(err)
	}
	err =exec.Command("sips","-z","128","128", full, "--out", target_folder+"/icon_64x64@2x.png").Run()
	if err!=nil{
		panic(err)
	}
	err =exec.Command("sips","-z","128","128", full, "--out", target_folder+"/icon_128x128.png").Run()
	if err!=nil{
		panic(err)
	}
	err =exec.Command("sips","-z","256","256", full, "--out", target_folder+"/icon_128x128@2x.png").Run()
	if err!=nil{
		panic(err)
	}
	err =exec.Command("sips","-z","256","256", full, "--out", target_folder+"/icon_256x256.png").Run()
	if err!=nil{
		panic(err)
	}
	err =exec.Command("sips","-z","512","512", full, "--out", target_folder+"/icon_256x256@2x.png").Run()
	if err!=nil{
		panic(err)
	}
	err =exec.Command("sips","-z","512","512", full, "--out", target_folder+"/icon_512x512.png").Run()
	if err!=nil{
		panic(err)
	}
	err =exec.Command("sips","-z","1024","1024", full, "--out", target_folder+"/icon_512x512@2x.png").Run()
	if err!=nil{
		panic(err)
	}
	err =exec.Command("iconutil","-c","icns",target_folder,"-o",target_folder+"/Icon.icns").Run()
	if err!=nil{
		panic(err)
	}
	cl.BoldGreen.Println("congratulations to you, Icon.icns was generated in your desktop....")
}