display_x = 165;
display_y = 100;

border_x = 20;
border_y = 15;

back_z = 3;
side_z = 28;

side = 3;

case_x = display_x + 2 * border_x;
case_y = display_y + 2 * border_y;

// Board: https://waveshare.com/compute-module-4-poe-board-b.htm
// These are only the holes:
board_x = 118.5;
board_y = 102.8;
board_left = case_x / 2 - board_x / 2;
board_bottom = case_y / 2 - board_y / 2;

// PoE LAN module
// Will be replaced by keystone modules later
lan_x = 80;
lan_y = 29;
lan_z = 23;


// ########### Back ###########

difference() {
    // back
    cube([display_x + 2 * border_x, display_y + 2 * border_y, back_z]);

    // vents (board)
    for (x = [0: 5: board_x]) {
        color("red")
        translate([x + board_left, board_bottom + 16, -1])
            cube([2, board_y - 32, 8]);
    }
}

// board screw points
for (x = [0, board_x], y = [0, board_y]) {
    translate([board_left + x, board_bottom + y, back_z])
        difference() {
            cylinder(8, d=6, $fn=25);
            cylinder(9, d=3, $fn=25);
        }
}

// ############ SIDE ########################


difference() {
    cube([display_x + 2 * border_x, display_y + 2 * border_y, side_z]);

    // cutout
    difference() {
        translate([side, side, -1])
            cube([display_x + 2 * border_x - 2 * side, display_y + 2 * border_y - 2 * side, side_z + 2]);
        // back plate screw holes
        for (x = [4, case_x - 4], y = [4, case_y - 4]) {
            translate([x, y, 0])
                cylinder(side_z, d=8, $fn=25);
        }
    }

    // PoE LAN module
    translate([-1, 8, side_z - lan_z + 0.1])
        cube([50, lan_y, lan_z]);

    // back plate screw holes
    for (x = [4, case_x - 4], y = [4, case_y - 4]) {
        color("red")
        translate([x, y, -0.1])
            cylinder(side_z + 2, d=3, $fn=25);
    }

    // vents
    for (x = [20:6:case_x-20]) {
        translate([x, -1, 4])
            cube([1, 10, side_z - 8]);
        translate([x, case_y-5, 4])
            cube([1, 10, side_z - 8]);
    }
}
