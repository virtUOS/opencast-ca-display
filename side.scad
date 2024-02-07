display_x = 165;
display_y = 100;
display_z = 28;

border_x = 20;
border_y = 15;

side = 3;

case_x = display_x + 2 * border_x;
case_y = display_y + 2 * border_y;

// Board: https://waveshare.com/w/upload/7/73/CM4-IO-BASE-A-details-size.jpg
// These are only the holes:
board_x = 85 - 3.5 - 23.5;
board_y = 56 - 2 * 3.5;

// PoE LAN module
// Will be replaced by keystone modules later
lan_y = 29;
lan_z = 23;

difference() {
    cube([display_x + 2 * border_x, display_y + 2 * border_y, display_z]);

    // cutout
    difference() {
        translate([side, side, -1])
            cube([display_x + 2 * border_x - 2 * side, display_y + 2 * border_y - 2 * side, display_z + 2]);
        // back plate screw holes
        for (x = [4, case_x - 4], y = [4, case_y - 4]) {
            translate([x, y, 0])
                cylinder(display_z, d=8, $fn=25);
        }
    }

    // PoE LAN module
    translate([-1, 8, display_z - lan_z + 0.1])
        cube([50, lan_y, lan_z]);

    // back plate screw holes
    for (x = [4, case_x - 4], y = [4, case_y - 4]) {
        color("red")
        translate([x, y, -0.1])
            cylinder(display_z + 2, d=3, $fn=25);
    }

    // vents
    for (x = [20:6:case_x-20]) {
        translate([x, -1, 4])
            cube([1, 10, display_z - 8]);
        translate([x, case_y-5, 4])
            cube([1, 10, display_z - 8]);
    }
}
