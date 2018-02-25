def trafficLight(light):
  if(light == "red"):
	  return("Stop")
  elif (light == "yellow"):
	  return("Stand By")
  elif(light == "green"):
	  return("Run")

def main():
  light = input("Please input the light(red, yellow, and green)");
  trafficLight(light);

if __name__ == "__main__": main()
