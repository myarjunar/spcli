/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/myarjunar/spcli/pkg/pathfinder"
	"github.com/spf13/cobra"
)

var (
	sp string
	pl string
)

var rootCmd = &cobra.Command{
	Use:   "spcli <graph.csv>",
	Short: "Find the sortest path between 2 nodes given graph in the csv file.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("you need to provide a graph.csv input")
		}

		filename := args[0]

		if sp != "" {
			pf := pathfinder.NewPathfinder(filename, sp)
			shortestDistance, err := pf.FindShortestPath()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(shortestDistance)
			}
		}

		if pl != "" {
			pf := pathfinder.NewPathfinder(filename, pl)
			totalDistance, err := pf.FindPathLength()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(totalDistance)
			}
		}

		return nil
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringVar(&sp, "sp", "", "find shortest path of 2 nodes e.g. AB")
	rootCmd.PersistentFlags().StringVar(&pl, "pl", "", "find path length of node sequence e.g. AEBC")
}
