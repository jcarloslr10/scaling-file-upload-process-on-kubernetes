import http from 'k6/http';
import { sleep } from 'k6';
import { htmlReport } from "https://raw.githubusercontent.com/benc-uk/k6-reporter/main/dist/bundle.js";
import { textSummary } from "https://jslib.k6.io/k6-summary/0.0.1/index.js";

export const options = {
    discardResponseBodies: true,
    scenarios: {
        fileapi: {
            executor: 'constant-vus',
            vus: 20,
            duration: '1m00s',
        },
    },
};

const file = open('sampleFile.pdf', 'b');

export default function () {
    const data = {
        file: http.file(file, 'sampleFile.pdf'),
    };

    http.post('http://fileapi.com/upload', data);
}

export function handleSummary(data) {
    return {
        "result.html": htmlReport(data),
        stdout: textSummary(data, { indent: " ", enableColors: true }),
    };
}