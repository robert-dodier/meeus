package main

import (
    "fmt"
    "math"

	"github.com/soniakeys/unit"
    "github.com/soniakeys/meeus/v3/sundial"
)

func print_line (name string, l sundial.Line) {

    fmt.Println (name, "@hour:", l.Hour, "$")
    fmt.Println (name, "@points: [")

    for i, xy := range l.Points {
        if i < len (l.Points) - 1 {
            fmt.Println ("Point (", xy.X, ",", xy.Y, "),")
        } else {
            fmt.Println ("Point (", xy.X, ",", xy.Y, ")")
        }
    }

    fmt.Println ("] $")
}

func print_line_with_declination (name string, l sundial.LineWithDeclination) {

    fmt.Println (name, "@declination:", l.Declination, "$")
    fmt.Println (fmt.Sprintf ("%s @description: \"%s\" $", name, l.Description))
    fmt.Println (name, "@points: [")

    for i, xy := range l.Points {
        if i < len (l.Points) - 1 {
            fmt.Println ("PointWithHour (", xy.H, ",", xy.X, ",", xy.Y, "),")
        } else {
            fmt.Println ("PointWithHour (", xy.H, ",", xy.X, ",", xy.Y, ")")
        }
    }

    fmt.Println ("] $")
}

func main () {

    my_latitude := unit.Angle (math.Pi/4)

    equatorial_lines_north, equatorial_lines_south  := sundial.Equatorial (my_latitude, 1.0)

    for i, l := range equatorial_lines_north {
        fmt.Println (fmt.Sprintf ("equatorial_lines_north[%d]: new (Line) $", i))
        print_line (fmt.Sprintf ("equatorial_lines_north[%d]", i), l)
    }

    for i, l := range equatorial_lines_south {
        fmt.Println (fmt.Sprintf ("equatorial_lines_south[%d]: new (Line) $", i))
        print_line (fmt.Sprintf ("equatorial_lines_south[%d]", i), l)
    }

    horizontal_lines, horizontal_center, horizontal_u := sundial.Horizontal (my_latitude, 1.0)
    fmt.Println ("horizontal_center: Point (", horizontal_center.X, ",", horizontal_center.Y, ") $")
    fmt.Println ("horizontal_u:", horizontal_u, "$")
    for i, l := range horizontal_lines {
        fmt.Println (fmt.Sprintf ("horizontal_lines[%d]: new (Line) $", i))
        print_line (fmt.Sprintf ("horizontal_lines[%d]", i), l)
    }

    horizontal_lines_by_declination, horizontal_center, horizontal_u := sundial.HorizontalByDeclination (my_latitude, 1.0)
    fmt.Println ("horizontal_center: Point (", horizontal_center.X, ",", horizontal_center.Y, ") $")
    fmt.Println ("horizontal_u:", horizontal_u, "$")
    for i, l := range horizontal_lines_by_declination {
        fmt.Println (fmt.Sprintf ("horizontal_lines_by_declination[%d]: new (LineWithDeclination) $", i))
        print_line_with_declination (fmt.Sprintf ("horizontal_lines_by_declination[%d]", i), l)
    }

    // assume sundial faces due south, so D = 0 if I understand correctly
    vertical_lines, vertical_center, vertical_u := sundial.Vertical (my_latitude, 0.0, 1.0)
    fmt.Println ("vertical_center: Point (", vertical_center.X, ",", vertical_center.Y, ") $")
    fmt.Println ("vertical_u:", vertical_u, "$")
    for i, l := range vertical_lines {
        fmt.Println (fmt.Sprintf ("vertical_lines[%d]: new (Line) $", i))
        print_line (fmt.Sprintf ("vertical_lines[%d]", i), l)
    }

}
