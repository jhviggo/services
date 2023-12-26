// k6 run -u 100 -d 5m ../infrastructure/k6-test.js

import http from "k6/http";
import { check } from "k6";

export const options = {
  iterations: 1,
};

export default function () {
  const response = http.get("https://master-service.api.viggo.work/health");
  check(response, {
    "is status 200": (r) => r.status === 200
  });
}