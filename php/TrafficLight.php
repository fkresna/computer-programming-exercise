<?php
	//PHP is usually made for web so it is a uncommon to accept input from command line
	echo "Please input the light(red, yellow, and green)";
	$light = fgets(STDIN);
	//Comparison in PHP usually use == operator but for string comparison it is better to use strcmp method
	if(strcmp($light,"red")) {
		echo "Stop";
	} else if (strcmp($light,"yellow")) {
		echo "Stand By";
	} else if (strcmp($light,"green")) {
		echo "Run";
	} 
?>