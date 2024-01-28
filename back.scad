display_x = 165;
display_y = 100;
display_z = 3;

border_x = 20;
border_y = 15;

case_x = display_x + 2 * border_x;
case_y = display_y + 2 * border_y;

// Board: https://waveshare.com/w/upload/7/73/CM4-IO-BASE-A-details-size.jpg
// These are only the holes:
board_x = 85 - 3.5 - 23.5;
board_y = 56 - 2 * 3.5;

// PoE LAN module
// Will be replaced by keystone modules later
lan_x = 80;
lan_y = 29;
lan_z = 23;

difference() {
    cube([display_x + 2 * border_x, display_y + 2 * border_y, display_z]);

    // back plate screw holes
    for (x = [4, case_x - 4], y = [4, case_y - 4]) {
        color("red")
        translate([x, y, -0.1])
            cylinder(10, d=4, $fn=25);
    }

    // vents
    for (x = [-board_x: 5: 0]) {
        color("red")
        translate([x + case_x - 50, case_y - board_y - 2, -1])
            cube([2, board_y - 16, 8]);
    }
}

// PoE LAN module holder
translate([case_x - lan_x - 20, 8, display_z])
    difference() {
        cube([20, lan_y, lan_z / 2]);
        translate([-1, 6, -1])
            cube([22, lan_y - 12, lan_z]);
        translate([10, lan_y / 2, lan_z / 4])
            rotate([90, 0, 0])
            cylinder(1.1 * lan_y, d=6, center=true, $fn=30);
    }

// board screw points
for (x = [-board_x, 0], y = [-board_y, 0]) {
    translate([x + case_x - 50, y + case_y - 10, display_z])
        difference() {
            cylinder(8, d=6, $fn=25);
            cylinder(9, d=3, $fn=25);
        }
}
