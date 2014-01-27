package main

import (
    "fmt"
    "github.com/codegangsta/cli"
    dockerclient "github.com/fsouza/go-dockerclient"
    "log"
    "os"
)

func main() {
    app := cli.NewApp()
    app.Name = "runpg"
    app.Usage = "Run a postgresql docker container"
    app.Flags = []cli.Flag{
        cli.StringFlag{"e, entrypoint", "unix:///var/run/docker.sock", "Entrypoint for docker's remote API"},
        cli.StringFlag{"p, pg-container", "cybersiddhu/centos-pgfdw:9.2", "Postgresql container image"},
        cli.StringFlag{"d, data-container", "cybersiddhu/pg-data:latest", "Data container image"},
    }
    app.Action = func(c *cli.Context) {
        err := RunContainer(c.String("entrypoint"), c.String("pg-container"), c.String("data-container"))
        if err != nil {
            log.Fatal(err)
        }
    }
    app.Run(os.Args)

}

func RunContainer(e, p, d string) error {
    client, err := dockerclient.NewClient(e)
    if err != nil {
        return err
    }

    copt := dockerclient.ListContainersOptions{All: true, Size: false}
    containers, err := client.ListContainers(copt)
    if err != nil {
        return err
    }

    //First check if there is any container listed
    //under that image name
    dcount := 0
    var dc string
    for _, c := range containers {
        if c.Image == d {
            dc = c.ID
            dcount += 1
        }
    }
    if dcount > 1 {
        return fmt.Errorf("Multiple data containers exists for %s image", d)
    } else if dcount == 1 { //data container exists
        //start the pg container

    } else { //data no container exist
        fmt.Println(dc)

    }
    return nil

}
