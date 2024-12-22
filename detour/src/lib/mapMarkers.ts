import type { Content, LatLngExpression } from 'leaflet';

export interface Marker {
  id: number;
  coords: LatLngExpression,
  contents?: Content;
}

interface Point {
  coords: LatLngExpression,
  hoverText: string;
  clickable?: boolean;
  articleId?: string;
}

export const markers: Marker[] = [
  {
    id: 1,
    coords: [56.156, 10.184],
    contents: '<b>The beginning</b><br>Aarhus where it all starts!',
  }
];

export const points: Point[] = [
  {
    coords: [56.156, 10.184],
    hoverText: 'Aarhus, Denmark',
    clickable: true,
    articleId: 't1'
  },
  {
    coords: [52.376, 9.750],
    hoverText: 'Hanover, Germany'
  },
  {
    coords: [48.129, 11.578],
    hoverText: 'Munich, Germany',
    clickable: true,
    articleId: 't2'
  },
  {
    coords: [45.464, 9.189],
    hoverText: 'Milan, Italy'
  },
  {
    coords: [44.493, 11.342],
    hoverText: 'Bologna, Italy'
  },
  {
    coords: [45.436, 12.333],
    hoverText: 'Venice, Italy'
  },
  {
    coords: [46.049, 14.507],
    hoverText: 'Ljubljana, Slovenia'
  },
  {
    coords: [47.06644, 15.43154],
    hoverText: 'Graz, Austria'
  },
  {
    coords: [47.4898, 19.0402],
    hoverText: 'Budapest, Hungary'
  },
  {
    coords: [45.7900, 17.3750],
    hoverText: 'Virovitica, Croatia'
  },
  {
    coords: [45.81136, 15.97748],
    hoverText: 'Zagreb, Croatia'
  },
  {
    coords: [44.11722, 15.24279],
    hoverText: 'Zadar, Croatia'
  },
  {
    coords: [43.51152, 16.43997],
    hoverText: 'Split, Croatia'
  },
  {
    coords: [43.34334, 17.80777],
    hoverText: 'Mostar, Bosnia'
  },
  {
    coords: [43.85262, 18.38530],
    hoverText: 'Sarajevo, Bosnia'
  },
  {
    coords: [44.81460, 20.46023],
    hoverText: 'Belgrade, Serbia'
  },
  {
    coords: [44.31821, 23.79601],
    hoverText: 'Craiova, Romania'
  },
  {
    coords: [44.43525, 26.10263],
    hoverText: 'Bucharest, Romania'
  },
  {
    coords: [43.20674, 27.91607],
    hoverText: 'Varna, Bulgaria'
  },
  {
    coords: [42.69708, 23.32172],
    hoverText: 'Sofia, Bulgaria'
  },
  {
    coords: [41.99583, 21.43149],
    hoverText: 'Skopje, Macedonia'
  },
  {
    coords: [41.32772, 19.81914],
    hoverText: 'Tirana, Albania'
  },
  {
    coords: [40.6385, 22.9379],
    hoverText: 'Thessaloniki, Greece'
  },
  {
    coords: [40.84739, 25.87503],
    hoverText: 'Alexandroupoli, Greece'
  },
  {
    coords: [41.0416, 28.9795],
    hoverText: 'Istanbul, Türkiye'
  },
  {
    coords: [39.9198, 32.8535],
    hoverText: 'Ankara, Türkiye'
  },
  {
    coords: [38.72195, 35.48711],
    hoverText: 'Kayseri, Türkiye'
  },
  {
    coords: [39.74666, 39.49111],
    hoverText: 'Erzincan, Türkiye'
  },
  {
    coords: [41.02447, 40.51962],
    hoverText: 'Rize, Türkiye'
  },
  {
    coords: [42.26994, 42.70513],
    hoverText: 'Kutaisi, Georgia'
  },
  {
    coords: [41.69772, 44.79983],
    hoverText: 'Tbilisi, Georgia'
  },
  {
    coords: [40.17481, 44.51434],
    hoverText: 'Yerevan, Armenia'
  },
  {
    coords: [40.68051, 46.35929],
    hoverText: 'Ganja, Azerbaijan'
  },
  {
    coords: [40.37376, 49.83310],
    hoverText: 'Baku, Azerbaijan'
  },
  {
    coords: [39.45535, 48.54599],
    hoverText: 'Bilasuvar, Azerbaijan'
  },
  {
    coords: [38.07405, 46.29637],
    hoverText: 'Tabriz, Iran'
  },
  {
    coords: [36.67937, 48.49759],
    hoverText: 'Zanjan, Iran'
  },
  {
    coords: [35.6885, 51.3895],
    hoverText: 'Tehran, Iran'
  },
  {
    coords: [33.98707, 51.44328],
    hoverText: 'Kashan, Iran'
  },
  {
    coords: [32.6695, 51.6645],
    hoverText: 'Isfahan, Iran'
  },
  {
    coords: [30.40072, 55.99368],
    hoverText: 'Rafsanjan, Iran'
  },
  {
    coords: [29.10639, 58.35761],
    hoverText: 'Bam, Iran'
  },
  {
    coords: [29.49792, 60.84419],
    hoverText: 'Zahedan, Iran'
  },
  {
    coords: [28.89639, 64.40955],
    hoverText: 'Dalbandin, Pakistan'
  },
  {
    coords: [30.1956, 67.0176],
    hoverText: 'Quetta, Pakistan'
  },
  {
    coords: [27.6988, 68.8560],
    hoverText: 'Sukkur, Pakistan'
  },
  {
    coords: [28.4186, 70.3048],
    hoverText: 'Rahim Yar Khan, Pakistan'
  },
  {
    coords: [30.1950, 71.4721],
    hoverText: 'Multan, Pakistan'
  },
  {
    coords: [31.5622, 74.3169],
    hoverText: 'Lahore, Pakistan'
  },
  {
    coords: [30.90768, 75.85164],
    hoverText: 'Ludhiana, India'
  },
  {
    coords: [28.6150, 77.2091],
    hoverText: 'New Delhi, India'
  },
  {
    coords: [26.8382, 80.9334],
    hoverText: 'Lucknow, India'
  },
  {
    coords: [26.7604, 83.3683],
    hoverText: 'Gorakhpur, India'
  },
  {
    coords: [28.2076, 83.9914],
    hoverText: 'Pokhara, Nepal'
  },
  {
    coords: [27.7073, 85.3204],
    hoverText: 'Kathmandu, Nepal'
  },
  {
    coords: [26.65045, 84.91151],
    hoverText: 'Motihari, India'
  },
  {
    coords: [24.7967, 85.0080],
    hoverText: 'Gaya, India'
  },
  {
    coords: [24.0418, 84.0697],
    hoverText: 'Medininagar, India'
  },
  {
    coords: [23.12197, 83.19790],
    hoverText: 'Ambikapur, India'
  },
  {
    coords: [22.08055, 82.14057],
    hoverText: 'Bilaspur, India'
  },
  {
    coords: [21.09588, 81.02946],
    hoverText: 'Rajnandgaon, India'
  },
  {
    coords: [21.1502, 79.0819],
    hoverText: 'Nagpur, India'
  },
  {
    coords: [19.67589, 78.53421],
    hoverText: 'Adilabad, India'
  },
  {
    coords: [17.36056, 78.47412],
    hoverText: 'Hyderabad, India'
  },
  {
    coords: [16.50902, 80.61868],
    hoverText: 'Vijayawada, India'
  },
  {
    coords: [14.44931, 79.98741],
    hoverText: 'Nellore, India'
  },
  {
    coords: [13.08207, 80.28172],
    hoverText: 'Chennai, India'
  }
];