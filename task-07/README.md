[PLEASE CLICK ME](http://akshaj000.github.io/2021/09/29/webscrapping/) 
ATTEMPT-1

I tried to scrap the data from forbes and i couldnt get the data on a csvfile as well as i couldnt print it. I tried the same on bloomberg page which is a similar website and it worked fine. Ive attached the program script and the csv file in test directory. Also ive attached the script to scrap from forbes here. I didnt include any csv file since its empty.

ATTEMPT-2

I tried goquery also, and i failed in scraping data from forbes. Ig i kinda gave up. Ig i tried my level best. I could have studied more about goquery, i will try to get back to this task at the end.
I tried goquery on other websites. It worked fine. I will attach my code, but its not the proper one to scrap the data to a csv, i just wanted to see whether its working with forbes elements.I will attach both the file try1 and try2.

please Kindly go through test directory in which i tried to scrap from bloomberg, which is a similar website like forbes using colly.

ATTEMPT-3

This time i figured out where i was wrong and understood html codes to forbes were java rendered. Geziyor was the frame which supported it. Its easy after identifying it.Its possible to fetch required part of the html to a json file using this framework, and is easy to convert it to a csv.But there was issue. Go couldnt identify  class="Country/Territory". Even i tried it with "//". It didnt help. I couldnt print the Country/Territory. I couldnt come up with a btr approach. How i did it was, I used geziyor to scrap all the html codes to a txt file, editted each "Country/Territory"  to "Country". Made a local web server and used colly to scrap it. This time it worked, But i had to run two files. One with the geziyor frame work which scraped the link and made a local server. Then one with colly which actually scraps the required elements and save it inside a csv file.I should have tried to find a way to include "/" and im still working on it.
Please follow this method to use my scraper:
1.Head to working_one directory
2.Run the gey-serve.go file inside the serve_here directory
3.Run the colly.go file inside the working_one directory

Code to run go file : go run . or go run filename.go

## ATTEMPT-2.5 (THE FINAL)

I actually managed to scrap directly using geyzior...i should have used ```'\\/'``` instead of ```/``` or ```//```.
The file pasted on this same directory works fine without any complication. This was actually an easy task which i made it harder.

