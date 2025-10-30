from flask import Flask, request, jsonify
from flask_sqlalchemy import SQLAlchemy
import os

import flask
print("Flask version:", flask.__version__)

app = Flask(__name__)
PORT = int(os.environ.get("PORT", 5000))

# PostgreSQL DB config
POSTGRES_USER = os.environ.get('POSTGRES_USER', 'postgres')
POSTGRES_PASSWORD = os.environ.get('POSTGRES_PASSWORD', 'postgres')
POSTGRES_DB = os.environ.get('POSTGRES_DB', 'coursesdb')
POSTGRES_HOST = os.environ.get('POSTGRES_HOST', 'postgres')
POSTGRES_PORT = os.environ.get('POSTGRES_PORT', '5432')

app.config['SQLALCHEMY_DATABASE_URI'] = f'postgresql://{POSTGRES_USER}:{POSTGRES_PASSWORD}@{POSTGRES_HOST}:{POSTGRES_PORT}/{POSTGRES_DB}'
app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False

db = SQLAlchemy(app)

# Course model
class Course(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    title = db.Column(db.String(100), nullable=False)
    instructor = db.Column(db.String(100), nullable=False)
    credits = db.Column(db.Integer, nullable=False)

    def to_dict(self):
        return {
            "id": self.id,
            "title": self.title,
            "instructor": self.instructor,
            "credits": self.credits
        }

# Routes

# health check
@app.route('/')
def home():
    return 'Course Service is running'

# create a course
@app.route('/courses', methods=['POST'])
def create_course():
    data = request.get_json()
    try:
        course = Course(
            title=data['title'],
            instructor=data['instructor'],
            credits=data['credits']
        )
        db.session.add(course)
        db.session.commit()
        return jsonify(course.to_dict()), 201
    except Exception as e:
        return jsonify({'error': 'Failed to create course', 'details': str(e)}), 500

# get a course via id
@app.route('/courses/<int:course_id>', methods=['GET'])
def get_course(course_id):
    course = Course.query.get(course_id)
    if not course:
        return jsonify({'error': 'Course not found'}), 404
    return jsonify(course.to_dict())

# get all courses
@app.route('/courses', methods=['GET'])
def get_all_courses():
    try:
        courses = Course.query.all()
        courses_list = [course.to_dict() for course in courses]
        return jsonify(courses_list)
    except Exception as e:
        return jsonify({'error': 'Failed to fetch courses', 'details': str(e)}), 500

if __name__ == '__main__':
    print("Flask version:", flask.__version__)
    print("App type:", type(app))

    with app.app_context():
        db.create_all()

    app.run(host='0.0.0.0', port=PORT)

