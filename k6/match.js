import http from 'k6/http';
import { check, sleep } from 'k6';
import { randomString, randomIntBetween, randomItem } from 'https://jslib.k6.io/k6-utils/1.1.0/index.js';


export let options = {
    stages: [
        { duration: '10s', target: 10000 },
        { duration: '20s', target: 10000 },
        { duration: '10s', target: 0 },
    ]
};

const BASE_URL = 'http://127.0.0.1:8080/api/v1/person';

function generateRandomName() {
    return `${randomString(25)}`;
}

function generateRandomPerson() {
    return {
        name: generateRandomName(),
        height: randomIntBetween(1, 1000),
        gender: randomItem(['male', 'female']),
        wanted_dates: randomIntBetween(1, 20)
    };
}

let addPersonHeaders = {
    headers: {
        'Content-Type': 'application/json',
    },
};


export default function () {
    let poll = [];

    // Add 2 people to the poll
    for (let i = 0; i < 2; i++) {
        let person = generateRandomPerson();
        let addPersonBody = JSON.stringify(person);
        let addPersonResponse = http.post(`${BASE_URL}/`, addPersonBody, addPersonHeaders);
        check(addPersonResponse, {
            'is status 200': (r) => r.status === 200,
        });

        poll.push(person.name);
    }

    // Remove a random person from  poll
    let randomIndex = Math.floor(Math.random() * poll.length);
    let nameToRemove = poll[randomIndex];
    let removePersonResponse = http.del(`${BASE_URL}/${nameToRemove}`);
    check(removePersonResponse, {
        'is status 200': (r) => r.status === 200,
    });
    poll.splice(randomIndex, 1);

    // Query people
    let queryPeopleResponse = http.get(`${BASE_URL}/?limit=${poll.length}`);
    check(queryPeopleResponse, {
        'is status 200': (r) => r.status === 200,
    });

    sleep(1);
}