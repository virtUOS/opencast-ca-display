display_x = 165.4;
display_y = 100.2;
display_z = 4;

border_x = 20 - 0.2;
border_y = 15 - 0.1;

case_x = display_x + 2 * border_x;
case_y = display_y + 2 * border_y;

difference() {
    cube([display_x + 2 * border_x, display_y + 2 * border_y, display_z]);
    // display cutout
    translate([border_x, border_y, -0.1])
        cube([display_x, display_y, display_z + 4.2]);

    // display clamp screw holes
    for (x = [-5, display_x + 5], y = [5, display_y - 5]) {
        color("red")
        translate([x + border_x, y + border_y, -0.1])
            cylinder(10, d=4, $fn=25);
    }

    // front plate screw holes
    for (x = [4, case_x - 4], y = [4, case_y - 4]) {
        color("red")
        translate([x, y, -0.1])
            cylinder(10, d=4, $fn=25);
    }
}


// mini display clip (to hold the top in place)
color("red")
translate([case_x / 2 - 20, border_y + display_y - 1, 0])
cube([40, 4, 1]);
