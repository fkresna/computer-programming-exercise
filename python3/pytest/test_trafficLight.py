import TrafficLight

def test_red():
  assert TrafficLight.trafficLight("red") == "Stop"
  assert TrafficLight.trafficLight("yellow") == "Stand By"
  assert TrafficLight.trafficLight("green") == "Run"
