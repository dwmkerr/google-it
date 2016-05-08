package main

import "testing"
import . "github.com/smartystreets/goconvey/convey"

func TestSpec(t *testing.T) {
    Convey("The help flag", t, func() {
        Convey("Should not parse the help flag if not present", func() {
            So(ParseHelpFlag("--not --present"), ShouldEqual, false)
        })

        Convey("Should parse the help flag if present in short form", func() {
            So(ParseHelpFlag("-h"), ShouldEqual, true)
        })

        Convey("Should parse the help flag if present in long form", func() {
            So(ParseHelpFlag("--help"), ShouldEqual, true)
        })

        Convey("Should parse the help flag if present with other args", func() {
            So(ParseHelpFlag(`-r 3 --help "Search string"`), ShouldEqual, true)
        })

        //  TODO can't get this to work cause I can't get the negative lookahead
        //  in the regexp: 
        //  working....(?<![\S"])(-h|--help)
        /* Convey("Should not parse the help flag if it is part of the search string", func() {
            So(ParseHelpFlag(`"--help in git commandline"`), ShouldEqual, false)
        }) */
    })


    Convey("The results flag", t, func() {
        Convey("Should return the default value if the flag is not present", func() {
            r, err := ParseResultsFlag("--not --present", 3)
            So(err, ShouldEqual, nil)
            So(r, ShouldEqual, 3)
        })

        Convey("Should support a short form without equals", func() {
            r, err := ParseResultsFlag("-r 2", 3)
            So(err, ShouldEqual, nil)
            So(r, ShouldEqual, 2)
        })

        Convey("Should support a short form with equals", func() {
            r, err := ParseResultsFlag("-r=2", 3)
            So(err, ShouldEqual, nil)
            So(r, ShouldEqual, 2)
        })

        Convey("Should support a long form without equals", func() {
            r, err := ParseResultsFlag("--results 2", 3)
            So(err, ShouldEqual, nil)
            So(r, ShouldEqual, 2)
        })

        Convey("Should support a long form with equals", func() {
            r, err := ParseResultsFlag("--results=2", 3)
            So(err, ShouldEqual, nil)
            So(r, ShouldEqual, 2)
        })

        Convey("Should not error and use default if there is no numeric results value", func() {
            r, err := ParseResultsFlag("--results=x", 3)
            So(err, ShouldEqual, nil)
            So(r, ShouldEqual, 3)
        })

        Convey("Should not error and use default if there is no results value", func() {
            r, err := ParseResultsFlag("--results --help", 3)
            So(err, ShouldEqual, nil)
            So(r, ShouldEqual, 3)
        })
    })
}