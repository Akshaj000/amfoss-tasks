import argparse
import requests
import webbrowser
import urllib
import sys

my_parser = argparse.ArgumentParser(description='Fetch Image from NASA API')
my_parser.add_argument('date',metavar='',help='Date in YYYY-MM-DD')
my_parser.add_argument('id',metavar='',type=int,help='Image ID')
my_parser.add_argument('-rovername',metavar='',help='Rover name - opportunity , spirit , curiosity(default)')
args=my_parser.parse_args()

rn=args.rovername

if str(rn) not in ['Curiosity','curiosity','Spirit','spirit','Opportunity','opportunity'] and rn!=None:
    print('unknown rovername')
    sys.exit()
elif rn == None:
    rovername='Curiosity'
else:
    rovername=str(rn)


url="https://api.nasa.gov/mars-photos/api/v1/rovers/{}/photos?earth_date={}&api_key=Oxnb3aAfbCBjXgreQ0EaFtKlPh66yNBl3QGWXgno".format(rovername,str(args.date))
data=requests.get(url)
id=args.id

data=data.json()

i=0
while True:
    try:
        if (data['photos'])[i]['id']==id:
            urllib.request.urlretrieve((data['photos'])[i]['img_src'],'newimage.jpg')
            break
        i+=1
    except:
        print('invalid parameter')
        break


