sudo apt update
sudo apt upgrade
sudo apt install xserver-xorg xinit x11-xserver-utils openbox chromium-browser

# Configure X.org server
sudo sed -i 's/^allowed_users=console/allowed_users=anybody/' /etc/X11/Xwrapper.config
cat ~/.xinitrc

# Install services
sudo cp browser.service /etc/systemd/system/
sudo cp startx.service /etc/systemd/system/
sudo systemctl start browser.service
sudo systemctl enable browser.service
sudo systemctl start startx.service
sudo systemctl enable startx.service

# TODO: opencast-ca-display.service
# Needs to run before browser service
