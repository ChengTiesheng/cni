package main

import (
	"os"

	"github.com/appc/cni/pkg/ns"
	"github.com/appc/cni/pkg/skel"
	"github.com/appc/cni/pkg/types"
	"github.com/vishvananda/netlink"
)

func cmdAdd(args *skel.CmdArgs) error {
	args.IfName = "lo" // ignore config, this only works for loopback
	err := ns.WithNetNSPath(args.Netns, false, func(hostNS *os.File) error {
		link, err := netlink.LinkByName(args.IfName)
		if err != nil {
			return err // not tested
		}

		err = netlink.LinkSetUp(link)
		if err != nil {
			return err // not tested
		}

		return nil
	})
	if err != nil {
		return err // not tested
	}

	result := types.Result{}
	return result.Print()
}

func cmdDel(args *skel.CmdArgs) error {
	args.IfName = "lo" // ignore config, this only works for loopback
	err := ns.WithNetNSPath(args.Netns, false, func(hostNS *os.File) error {
		link, err := netlink.LinkByName(args.IfName)
		if err != nil {
			return err // not tested
		}

		err = netlink.LinkSetDown(link)
		if err != nil {
			return err // not tested
		}

		return nil
	})
	if err != nil {
		return err // not tested
	}

	return nil
}

func main() {
	skel.PluginMain(cmdAdd, cmdDel)
}
