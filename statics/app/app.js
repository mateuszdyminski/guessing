'use strict';

var G = angular.module('G', ['ngRoute', 'ui.bootstrap'])

    .config(function($routeProvider) {
        $routeProvider
            .when('/game', {
                templateUrl: 'app/views/game.html',
                controller: 'GameCtrl'
            })
            .when('/game/:round', {
                templateUrl: 'app/views/game.html',
                controller: 'GameCtrl'
            })
            .otherwise({
                redirectTo: '/game'
            });
    });