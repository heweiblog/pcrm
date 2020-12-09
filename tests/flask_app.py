from flask import Flask,request,jsonify
app = Flask(__name__)

@app.route('/',methods=['GET', 'POST'])
def index():
	data = request.get_json()
	print(data)
	return jsonify({'status':'success','msg':''})
	    
if __name__ == '__main__':
	app.run(host='::',port=22222,debug=True)
