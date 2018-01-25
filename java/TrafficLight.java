import java.util.Scanner;

 class TrafficLight {
 	public static void main(String[] args) {
 		Scanner input = new Scanner(System.in);
 		System.out.print("Please input the light(red, yellow, and green)");
 		String light = input.next();
 		if (light.equals("red")) {
 			System.out.println("Stop");
 		} else if (light.equals("yellow")) {
 			System.out.println("Stand By");
 		} else if (light.equals("green")) {
 			System.out.println("Run");
 		}
 	}
 }
