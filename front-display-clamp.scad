display_x = 165;
display_y = 100 - 4;
display_z = 3;

border_x = 10;
border_y = 6;

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
            cylinder(10, d=2.9, $fn=25);
    }

    // top/down cutout
    translate([border_x, -0.1, display_z - 1.5])
        cube([display_x, case_y + 0.2, display_z + 4.2]);

    // cable cutout
    translate([3 * border_x, border_y + 10, -0.1])
        cube([display_x, display_y - 20, display_z + 4.2]);
}
