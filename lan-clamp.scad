border_x = 20;
border_y = 15;

back_z = 3;
side_z = 28;

side = 3;

case_x = 165 + 2 * border_x;
case_y = 100 + 2 * border_y;

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

// LAN cable clamp
difference() {
    for (x = [0:8 + lan_cable_d]) {
        translate([x, 0, 0])
            cylinder(3, d=8, $fn=25);
    }
    for (x = [0, 8 + lan_cable_d]) {
        translate([x, 0, -0.1])
            cylinder(lan_cable_d, d=3.6, $fn=25);
    }
}
