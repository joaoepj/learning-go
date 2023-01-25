// learning-go is a repository showing how to apply development
// best practices in your code while learning a programming language
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joaoepj/learning-go/lg_misc"
	problems "github.com/joaoepj/learning-go/lg_problems"
	"github.com/urfave/cli/v2"
)

var (
	LearningGoGitHash   = "None"
	LearningGoGitBranch = "None"
)

func main() {

	log.SetFlags(log.Lshortfile)

	app := &cli.App{
		Name:  "learning-go",
		Usage: "a repository to learn the go programming language.",
		UsageText: `learning-go command [command options]
learning-go [global options]`,
		HideVersion: false,
		Version:     LearningGoGitBranch + "-" + LearningGoGitHash,
		Flags:       []cli.Flag{},
		Action: func(c *cli.Context) error {
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "problems",
				Aliases: []string{"p"},
				Usage:   "Some HackerRank problem solutions implemented in golang.",
				Subcommands: []*cli.Command{
					{
						Name:    "bigsum",
						Aliases: []string{"bs"},
						Usage:   "A Very Big Sum Problem",
						Action: func(c *cli.Context) error {
							problems.AVeryBigSum([]int64{1.0, 2.0, 3.0})
							return nil
						},
					},
					{
						Name:    "catmouse",
						Aliases: []string{"cm"},
						Usage:   "Cats and a Mouse Problem",
						Action: func(c *cli.Context) error {
							problems.CatAndMouse(5, 9, 13)
							return nil
						},
					},
					{
						Name:    "climbinglearderboard",
						Aliases: []string{"cl"},
						Usage:   "Climbing the Leaderboard Problem",
						Action: func(c *cli.Context) error {
							problems.ClimbingLeaderboard2([]int32{100, 100, 50, 40, 40, 20, 10}, []int32{5, 25, 50, 120})
							return nil
						},
					},
					{
						Name:    "comparethetriplets",
						Aliases: []string{"ctt"},
						Usage:   "Compare The Triplets Problem",
						Action: func(c *cli.Context) error {
							t1 := []int32{1, 2, 3}
							t2 := []int32{1, 2, 3}
							problems.CompareTriplets(t1, t2)
							return nil
						},
					},
					{
						Name:    "designerpdfviewer",
						Aliases: []string{"dpv"},
						Usage:   "Designer Pdf Viewer Problem",
						Action: func(c *cli.Context) error {

							problems.DesignerPdfViewer(problems.Alphabet_height, "reality")

							return nil
						},
					},
					{
						Name:    "formingmagicsquare",
						Aliases: []string{"fms"},
						Usage:   "Forming A Magic Square Problem",
						Action: func(c *cli.Context) error {
							problems.FormingMagicSquare(problems.TheSquare)
							return nil
						},
					},
					{
						Name:    "hourglasssum",
						Aliases: []string{"hgs"},
						Usage:   "Hourglass Sum Problem",
						Action: func(c *cli.Context) error {
							problems.HourglassSum(problems.HourglassSum_arr)
							return nil
						},
					},
					{
						Name:    "hurdlerace",
						Aliases: []string{"hr"},
						Usage:   "Hurdle Race Problem",
						Action: func(c *cli.Context) error {
							problems.Staircase(5)
							return nil
						},
					},
					{
						Name:    "letfrotation",
						Aliases: []string{"lr"},
						Usage:   "Left Rotation Problem",
						Action: func(c *cli.Context) error {
							problems.RotLeft([]int32{1, 2, 3, 4, 5}, 3)
							return nil
						},
					},
					{
						Name:    "migratorybirds",
						Aliases: []string{"mb"},
						Usage:   "Migratory Birds Problem",
						Action: func(c *cli.Context) error {
							problems.MigratoryBirds(problems.MigratoryBirds_arr)
							return nil
						},
					},
					{
						Name:    "minimumbribes",
						Aliases: []string{"mb2"},
						Usage:   "Minimun Bribes Problem",
						Action: func(c *cli.Context) error {
							for _, arr := range problems.MinimumBribes_arr {
								problems.MinimumBribes(arr)
							}
							return nil
						},
					},
					{
						Name:    "minimumswaps2",
						Aliases: []string{"ms2"},
						Usage:   "Minimun Swaps 2 Problem",
						Action: func(c *cli.Context) error {
							for _, arr := range problems.MinimumSwaps2_arr {
								problems.MinimumSwaps2(arr)
							}
							return nil
						},
					},

					{
						Name:    "minmaxsum",
						Aliases: []string{"mms"},
						Usage:   "Min-Max Sum Problem",
						Action: func(c *cli.Context) error {
							arr := []int32{2, 8, 5, 9, 2, 5, 7, 1}
							problems.MinMaxSum(arr)
							return nil
						},
					},
					{
						Name:    "missinginteger",
						Aliases: []string{"mi"},
						Usage:   "Missing Integer Problem",
						Action: func(c *cli.Context) error {
							problems.TestMissingInteger()
							return nil
						},
					},
					{
						Name:    "staircase",
						Aliases: []string{"s"},
						Usage:   "Staircase Problem",
						Action: func(c *cli.Context) error {
							problems.Staircase(5)
							return nil
						},
					},
				}, // Closing Subcommands
			},
			{
				Name:    "misc",
				Aliases: []string{"m"},
				Usage:   "Miscelaneous code",
				Subcommands: []*cli.Command{
					{
						Name:    "asn1",
						Aliases: []string{"as"},
						Usage:   "Test encoding/decoding time with ASN.1 library",
						Action: func(c *cli.Context) error {
							lg_misc.Asn1()
							return nil
						},
					},
					{
						Name:    "mediansortedarrays",
						Aliases: []string{"msa"},
						Usage:   "Find Median Sorted Arrays",
						Action: func(c *cli.Context) error {
							x := []int{1, 2}
							y := []int{3, 4}
							fmt.Println("v1: ", lg_misc.FindMedianSortedArraysV1(x, y))
							fmt.Println("v2: ", lg_misc.FindMedianSortedArraysV2(x, y))
							fmt.Println("v3: ", lg_misc.FindMedianSortedArraysV3(x, y))
							return nil
						},
					},
					{
						Name:    "snmpget",
						Aliases: []string{"sg"},
						Usage:   "snmpget",
						Action: func(c *cli.Context) error {
							lg_misc.SnmpGet()
							return nil
						},
					},
					{
						Name:    "snmpwalk",
						Aliases: []string{"sw"},
						Usage:   "snmpwalt",
						Action: func(c *cli.Context) error {
							lg_misc.SnmpWalk()
							return nil
						},
					},
					{
						Name:    "testfdlimit",
						Aliases: []string{"tfl"},
						Usage:   "Testing the Limit for File Descriptors",
						Action: func(c *cli.Context) error {
							lg_misc.TestFdLimit()
							return nil
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
