package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(vectorCmd)
	vectorCmd.AddCommand(vectorCreate)
	vectorCmd.AddCommand(vectorDelete)
	vectorCmd.AddCommand(vectorLoad)
	vectorCmd.AddCommand(vectorStore)
	vectorCmd.AddCommand(vectorList)
}

var vectorCmd = &cobra.Command{
	Use:   "vector [command]",
	Short: "vector management",
	Long:  "Create, delete, load, store, list, and query vectors in the device",
	Run: func(cobra *cobra.Command, args []string) {
		fmt.Println("vector command")
	},
}

var vectorCreate = &cobra.Command{
	Use:   "create",
	Short: "create a new vector",
	Long:  "create a new vector allocation on the device",
	Run:   createVector,
}

func createVector(cobra *cobra.Command, args []string) {
	fmt.Println("vector create")
}

var vectorDelete = &cobra.Command{
	Use:   "delete",
	Short: "delete an existing vector",
	Long:  "delete a vector allocation on the device",
	Run:   deleteVector,
}

func deleteVector(cobra *cobra.Command, args []string) {
	fmt.Println("vector delete")
}

var vectorLoad = &cobra.Command{
	Use:   "load",
	Short: "load a vector into device memory",
	Long:  "load the content of a host vector into a previously allocated vector on the device",
	Run:   loadVector,
}

func loadVector(cobra *cobra.Command, args []string) {
	fmt.Println("vector load")
}

var vectorStore = &cobra.Command{
	Use:   "store",
	Short: "store a vector from device memory",
	Long:  "store an existing vector on the device in host memory",
	Run:   storeVector,
}

func storeVector(cobra *cobra.Command, args []string) {
	fmt.Println("vector store")
}

var vectorList = &cobra.Command{
	Use:   "list",
	Short: "list all vectors in device memory",
	Long:  "list all vectors allocated in device memory",
	Run:   listVector,
}

func listVector(cobra *cobra.Command, args []string) {
	fmt.Println("list vectors")
}
