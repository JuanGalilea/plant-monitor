#include <DHT_U.h>
#include <DHT.h>

#define DHTPIN 2
#define DHTTYPE DHT11

#define PIN_LDR A5

DHT dht(DHTPIN, DHTTYPE);

void setup()
{
  Serial.begin(9600);
  dht.begin();
}

void loop()
{
  delay(1000);
  float h = dht.readHumidity();
  float t = dht.readTemperature();
  int sensorValue = analogRead(PIN_LDR);

  float hic = dht.computeHeatIndex(t, h, false);

  if (isnan(h) || isnan(t))
  {
    Serial.println("Failed to read from DHT sensor!");
    return;
  }

  Serial.print(h);
  Serial.print(",");
  Serial.print(hic);
  Serial.print(",");
  Serial.println(sensorValue);

}
