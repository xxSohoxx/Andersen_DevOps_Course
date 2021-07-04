# This is a simple python app which receives JSON objects and return strings

# import main Flask class, request object. And from emoji importing emojize method
from flask import Flask, request
from emoji import emojize

# create the Flask app
app = Flask(__name__)

# set the default route

@app.route('/', methods=['POST'])
#function which receives data and returns strings
def json_example():
    request_data = request.get_json(force=True)

    reply = ""

    animal = request_data['animal']
    sound = request_data['sound']
    count = request_data['count']
    for i in range(count):
         reply+= (emojize(":" + animal + ":") + " says " + sound + "\n")
    reply+= "Made with Love by june!\n"
    return reply

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=80, debug=True)
