#!/bin/bash

DIR="/home/dongri/Pictures/wallpapers"
PIC=$(ls $DIR/* | shuf -n1)
gsettings set org.gnome.desktop.background picture-uri "file://$PIC"

