'use strict';

G.service('ResultsService', ['$http',
    function($http) {
        this.saveResult = function(data) {
            return $http({
                url: '/restapi/results',
                method: "POST",
                data: data
            });
        };
    }
]);