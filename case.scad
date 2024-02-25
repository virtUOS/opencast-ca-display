back_z = 3;
side_z = 32;

side = 3;

case_x = 205;
case_y = 132;

// Display: https://waveshare.com/7inch-hdmi-lcd-c.htm
// Display screw holes
display_x = 156.9;
display_y = 114.96;
display_z = 7.2;
// move display right by 1mm since the right bezel is 2mm wider
display_left = case_x / 2 - display_x / 2 + 1;
display_bottom = case_y / 2 - display_y / 2;

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

// https://amazon.de/dp/B07ZKHB72D
keystone_y = 23;
keystone_z = 17;
keystone_border_y = 41;
keystone_border_z = 19.3;

// LAN cable
lan_cable_d = 5.6;

// ########### Back ###########

difference() {
    // back
    cube([case_x, case_y, back_z]);

    // vents (board)
    for (x = [0: 5: board_x]) {
        color("red")
        translate([x + board_left, board_bottom + 16, -1])
            cube([2, board_y - 32, 8]);
    }

    // board screw point in back plate
    for (x = [0, board_x], y = [0, board_y]) {
        translate([board_left + x, board_bottom + y, 1])
            color("red")
            cylinder(9, d=3, $fn=25);
    }

    // wall mount screw holes
    for (x = [12, case_x - 12]) {
        translate([x, case_y - side - 5, -0.1])
            color("red")
            cylinder(back_z + 2, d=4, $fn=30);
    }
}

// board screw points
for (x = [0, board_x], y = [0, board_y]) {
    translate([board_left + x, board_bottom + y, back_z]) {
        difference() {
            cylinder(2, d=6, $fn=25);
            translate([0, 0, -1])
            color("red")
            cylinder(9, d=3, $fn=25);
        }
        if ((x <= 0 && y > 0) || (x > 0 && y <= 0))
            color("blue")
            cylinder(6, d=2.6, $fn=25);
    }
}

// display screw points
for (x = [0, display_x], y = [-1, 1]) {
    height = side_z - back_z - display_z;
    translate([display_left + x, display_bottom + (1 + y) * display_y / 2, back_z])
        difference() {
            cylinder(height, d=6, $fn=25);
            color("red")
            cylinder(height + 2, d=3, $fn=25);
        }

    // Pins to hold display in place
    if ((x <= 0 && y >= 0) || (x > 0 && y < 0))
        translate([display_left + x, display_bottom + (1 + y) * display_y / 2, back_z])
        color("blue")
        cylinder(height + 4, d=2.8, $fn=25);

    // connect screw points and side
    translate([display_left + x, display_bottom + (1 + y) * display_y / 2 + y * 5, height / 2 + back_z])
        cube([3, 5, height], center=true);
}

// wall mount
for (x = [12, case_x - 12]) {
    translate([x, case_y - side - 5, 0])
        difference() {
            cylinder(back_z + 1, d=10, $fn=30);
            color("red")
            cylinder(back_z + 2, d=4, $fn=30);
        }
}

// ############ SIDE ########################


difference() {
    cube([case_x, case_y, side_z]);

    // cutout
    difference() {
        translate([side, side, -1])
            cube([case_x - 2 * side, case_y - 2 * side, side_z + 2]);

        // front plate screw in cylinder
        for (x = [4, case_x - 4], y = [4, case_y - 4]) {
            translate([x, y, 0])
                cylinder(side_z, d=8, $fn=25);
        }
    }

    // front plate screw holes
    for (x = [4, case_x - 4], y = [4, case_y - 4]) {
        color("red")
        translate([x, y, -0.1])
            cylinder(side_z + 2, d=3, $fn=25);
    }

    // vents
    for (x = [42:6:case_x-42]) {
        translate([x, -1, 4])
            cube([1, 10, side_z - 8]);
        translate([x, case_y-5, 4])
            cube([1, 10, side_z - 8]);
    }

    /*
    // keystone module
    translate([-2, 10, back_z])
    union() {
        color("green")
        cube([3, keystone_border_y, keystone_border_z]);
        translate([0, (keystone_border_y - keystone_y) / 2, (keystone_border_z - keystone_z) / 2])
            color("red")
            cube([30, keystone_y, keystone_z]);
    }
    */

    // LAN cable cutout
    for (z = [0:side_z]) {
        translate([18, side + 1, back_z + lan_cable_d / 2 + z])
            color("red")
            rotate([90, 0, 0])
            cylinder(side + 2, d=lan_cable_d, $fn=25);
    }

    // LAN cable cutout
    for (z = [0:side_z]) {
        translate([-1, 14, back_z + lan_cable_d / 2 + z])
            color("red")
            rotate([0, 90, 0])
            cylinder(side + 2, d=lan_cable_d, $fn=25);
    }
}

// better structural integrity near LAN cable cutout
translate([0, 20, 0])
    cube([8, 4, side_z]);


// LAN cable clamp
translate([10, 45, back_z])
for (x = [0, 8 + lan_cable_d]) {
    translate([x, 0, 0])
    difference() {
        cylinder(lan_cable_d - 1, d=8, $fn=25);
        color("red")
        cylinder(lan_cable_d, d=3, $fn=25);
    }
}
