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
display_left = case_x / 2 - display_x / 2;
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

// display border
display_border_left = 25;
display_border_right = 25;
display_border_bottom = 26;
display_border_top = 18;
display_cutout_x = case_x - display_border_left - display_border_right;
display_cutout_y = case_y - display_border_top - display_border_bottom;

// ########### Front ###########

//projection()
difference() {
    cube([case_x, case_y, 1.5]);

    // display cutout
    translate([display_border_left, display_border_bottom, -0.1])
        cube([display_cutout_x, display_cutout_y, 9]);

    // front plate screw holes
    for (x = [4, case_x - 4], y = [4, case_y - 4]) {
        color("red")
        translate([x, y, -0.1])
            cylinder(10, d=4, $fn=25);
    }
}
