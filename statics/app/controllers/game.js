'use strict';

angular.module('G').controller('GameCtrl', function($scope, $location, $timeout, $routeParams, ResultsService, CFG) {
    $scope.state = {}
    $scope.step = 1;
    $scope.round = 1;
    $scope.images = [];
    $scope.img = '';
    $scope.index = 1;
    $scope.maxRounds = Object.keys(CFG.rounds).length;

    if($routeParams.round) {
        $scope.round = $routeParams.round;
    }

    $scope.startGame = function(round) {
        if (round) {
            $scope.round = round;
        }
        $scope.step = 2;

        $scope.images = CFG.rounds[$scope.round].images;
        var noOfImages = CFG.rounds[$scope.round].images.length;

        $scope.state.round = $scope.round;
        $scope.state.steps = [];
        $scope.showImg(0, noOfImages);
    };

    $scope.showImg = function(index, noOfImages) {
        if (index >= CFG.rounds[$scope.round].steps) {
            $scope.chooseAnswer();
            return
        }

        $scope.index = index + 1;

        var imageId = getRandomInt(0, noOfImages - 1)
        $scope.state.steps.push(imageId);
        $scope.img = CFG.rounds[$scope.round].images[imageId];
        $timeout(function() {$scope.showImg(index + 1, noOfImages)}, CFG.rounds[$scope.round].sleep);
    }

    $scope.chooseAnswer = function() {
        $scope.step = 3;
    }

    $scope.endRound = function(index) {
        $scope.state.answer = index;

        ResultsService.saveResult($scope.state)
        .success(function(response, status, headers) {
            // ask for next round
            $scope.step = 4;
        })
        .error(function(response) {
            console.log(response);
        });
    }

    function getRandomInt(min, max) {
        return Math.floor(Math.random() * (max - min + 1)) + min;
    }
});