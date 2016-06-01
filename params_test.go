package main

import "testing"
import . "github.com/smartystreets/goconvey/convey"

func TestSpec(t *testing.T) {

    Convey("The param loader", t, func() {

        Convey("Should handle no params", func() {
            params, err := ParseParams([]string{})
            So(params.ShowHelp.Present, ShouldEqual, false)
            So(params.Results.Present, ShouldEqual, false)
            So(params.Open.Present, ShouldEqual, false)
            So(err, ShouldEqual, nil)
        })

        Convey("Should handle help", func() {
            params, err := ParseParams([]string{"--help"})
            So(params.ShowHelp.Present, ShouldEqual, true)
            So(params.Results.Present, ShouldEqual, false)
            So(params.Open.Present, ShouldEqual, false)
            So(err, ShouldEqual, nil)
        })

        Convey("Should handle results", func() {
            params, err := ParseParams([]string{`-r=3 "Example"`})
            So(params.ShowHelp.Present, ShouldEqual, false)
            So(params.Results.Present, ShouldEqual, true)
            So(params.Results.Count, ShouldEqual, 3)
            So(params.Open.Present, ShouldEqual, false)
            So(err, ShouldEqual, nil)
        })

        Convey("Should handle open", func() {
            params, err := ParseParams([]string{`--open 5`})
            So(params.ShowHelp.Present, ShouldEqual, false)
            So(params.Results.Present, ShouldEqual, false)
            So(params.Open.Present, ShouldEqual, true)
            So(params.Open.LinkId, ShouldEqual, "5")
            So(err, ShouldEqual, nil)
        })
    })

    Convey("The help param", t, func() {

        Convey("Should not parse the help param if not present", func() {
            So(ParseHelpFlag("--not --present"), ShouldEqual, false)
        })

        Convey("Should parse the help param if present in short form", func() {
            So(ParseHelpFlag("-h"), ShouldEqual, true)
        })

        Convey("Should parse the help param if present in long form", func() {
            So(ParseHelpFlag("--help"), ShouldEqual, true)
        })

        Convey("Should parse the help param if present with other args", func() {
            So(ParseHelpFlag(`-r 3 --help "Search string"`), ShouldEqual, true)
        })

        //  TODO can't get this to work cause I can't get the negative lookahead
        //  in the regexp: 
        //  working....(?<![\S"])(-h|--help)
        /* Convey("Should not parse the help flag if it is part of the search string", func() {
            So(ParseHelpFlag(`"--help in git commandline"`), ShouldEqual, false)
        }) */
    })

    Convey("The results param", t, func() {

        Convey("Should return 'not present' if the param is not present", func() {
            present, _, err := ParseResultsFlag("--not --present")
            So(present, ShouldEqual, false)
            So(err, ShouldEqual, nil)
        })

        Convey("Should support a short form without equals", func() {
            present, r, err := ParseResultsFlag("-r 2")
            So(present, ShouldEqual, true)
            So(r, ShouldEqual, 2)
            So(err, ShouldEqual, nil)
        })

        Convey("Should support a short form with equals", func() {
            present, r, err := ParseResultsFlag("-r=2")
            So(present, ShouldEqual, true)
            So(r, ShouldEqual, 2)
            So(err, ShouldEqual, nil)
        })

        Convey("Should support a long form without equals", func() {
            present, r, err := ParseResultsFlag("--results 2")
            So(present, ShouldEqual, true)
            So(r, ShouldEqual, 2)
            So(err, ShouldEqual, nil)
        })

        Convey("Should support a long form with equals", func() {
            present, r, err := ParseResultsFlag("--results=2")
            So(present, ShouldEqual, true)
            So(r, ShouldEqual, 2)
            So(err, ShouldEqual, nil)
        })

        Convey("Should not error and return 'not present' if there is no results value", func() {
            present, _, err := ParseResultsFlag("--results --help")
            So(present, ShouldEqual, false)
            So(err, ShouldEqual, nil)
        })
    })

    Convey("The open param", t, func() {

        Convey("Should return 'not present' if the param is not present", func() {
            present, _ := ParseOpenFlag("--not --present")
            So(present, ShouldEqual, false)
        })

        Convey("Should support a short form without equals", func() {
            present, r := ParseOpenFlag("-o 1")
            So(present, ShouldEqual, true)
            So(r, ShouldEqual, "1")
        })

        Convey("Should support a short form with equals", func() {
            present, r := ParseOpenFlag("-o=2")
            So(present, ShouldEqual, true)
            So(r, ShouldEqual, "2")
        })

        Convey("Should support a long form without equals", func() {
            present, r := ParseOpenFlag("--open 3")
            So(present, ShouldEqual, true)
            So(r, ShouldEqual, "3")
        })

        Convey("Should support a long form with equals", func() {
            present, r := ParseOpenFlag("--open=4")
            So(present, ShouldEqual, true)
            So(r, ShouldEqual, "4")
        })

        Convey("Should return 'not present' if there is no open value", func() {
            present, _ := ParseOpenFlag("--open --help")
            So(present, ShouldEqual, false)
        })
    })
}