'use strict';

angular.module('G').controller('MenuCtrl', function($scope, $location) {
    $scope.isActive = function(viewLocation) {
        return viewLocation === $location.path();
    };
});