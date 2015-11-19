'use strict';

angular.module('G').constant('CFG', {
    'rounds': {
        '1': {
            'images': [ 'data/set1/1.jpg', 'data/set1/2.jpg' ],
            'image_width': 400,
            'steps': 8,
            'sleep': 1000
        },
        '2': {
            'images': [ 'data/set2/1.jpg', 'data/set2/2.jpg', 'data/set2/3.jpg' ],
            'image_width': 400,
            'steps': 10,
            'sleep': 1000
        }
    }
});