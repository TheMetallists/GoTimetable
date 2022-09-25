class ApiHelper {
    static loadTimeTable(result) {
        const formData = new FormData();
        fetch(process.env.REACT_APP_SERVER + '/api/telek', {
            method: "POST",
            body: formData
        })
            .then(res => res.json())
            .then(oData => {
                if (oData.error) {
                    // TODO: replace alert with something non blocking
                    alert("API LOADING ERROR!");
                    return;
                }

                result(oData.klassen, oData.hash);
            });


        /*let $arReturn = {};
        [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11].map($klassParallel => {
            ['А', 'Б', 'В', 'Г'].map($klassLetter => {
                $arReturn[$klassParallel + $klassLetter] = [];
                [1, 2, 3, 4, 5, 6].map($weekday => {
                    $arReturn[$klassParallel + $klassLetter].push({
                        day: $weekday,
                        lessons: ApiHelper.shuffle([
                            {lesson: 'Матан', cab: 4},
                            {lesson: 'Шизика', cab: 42},
                            {lesson: 'Химия', cab: 16},
                            {lesson: 'Биология', cab: 8},
                            {lesson: 'ОБЖорка', cab: 52},
                            {lesson: 'Музло', cab: 27},
                            {lesson: 'Русичи', cab: 49},
                            {lesson: 'Литерча', cab: 37},
                            {lesson: 'Физуха', cab: 38},
                            {lesson: 'Труд', cab: -1},
                            {lesson: 'English', cab: 44},
                            {lesson: 'Инф-ра', cab: '1111А'},
                            {lesson: 'Инф-ра', cab: '1111А'},
                            {lesson: 'Инф-ра', cab: '1111А'},
                        ]).slice(0, Math.floor(Math.random() * 8) + 4),
                        notice: ''
                    });
                })
            })
        });

        console.log(JSON.stringify($arReturn));

        result($arReturn);*/
    }

    static callApiUnauthorized(url, form, result) {
        fetch(process.env.REACT_APP_SERVER + url, {
            method: "POST",
            body: form
        })

            .then(res => {
                if (!res.ok) {
                    throw Error(res.statusText);
                }
                return res.json();
            })
            .then(oData => {
                if (oData.error) {
                    // TODO: replace alert with something non blocking
                    alert("API ERROR: " + oData.ermes);
                    return;
                }

                result(oData);
            })
            .catch(err => {
                alert("Api error 1: " + JSON.stringify(err));
            });
    }

    static token = '';

    static callApiAuthorized(url, form, result, __error) {
        let options = {
            method: "POST",
            headers: new Headers({
                'Authorization': 'Bearer ' + ApiHelper.token,
            }),
        };
        if (form) {
            options['body'] = form
        }

        fetch(process.env.REACT_APP_SERVER + url, options)
            .then(res => {
                if (!res.ok) {
                    throw Error(res.statusText);
                }
                return res.json();
            })
            .then(oData => {
                if (oData.error) {
                    // TODO: replace alert with something non blocking
                    if (typeof __error === 'function') {
                        __error(oData);
                    } else {
                        alert("API ERROR: " + oData.ermes);
                    }
                    return;
                }
                result(oData);
            })
            .catch(err => {
                console.log("API ERROR: ",err);
                if (typeof __error === 'function') {
                    __error(err, 2);
                } else {
                    alert("Api error 2: " + JSON.stringify(err));
                }
            });
    }

    static callApiAuthorizedJSON(url, _json, result, __error) {
        let options = {
            method: "POST",
            headers: new Headers({
                'Authorization': 'Bearer ' + ApiHelper.token,
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            }),
        };
        if (_json) {
            options['body'] = _json
        }

        fetch(process.env.REACT_APP_SERVER + url, options)
            .then(res => {
                if (!res.ok) {
                    throw Error(res.statusText);
                }
                return res.json();
            })
            .then(oData => {
                if (oData.error) {
                    // TODO: replace alert with something non blocking
                    if (typeof __error === 'function') {
                        __error(oData);
                    } else {
                        alert("API ERROR: " + oData.ermes);
                    }

                    return;
                }

                result(oData);
            })
            .catch(err => {
                if (typeof __error === 'function') {
                    __error(err, 2);
                } else {
                    console.log(err);
                    alert("Api error 2: " + JSON.stringify(err));
                }
            });
    }

    static shuffle(a) {
        for (let i = a.length - 1; i > 0; i--) {
            const j = Math.floor(Math.random() * (i + 1));
            [a[i], a[j]] = [a[j], a[i]];
        }
        return a;
    }
}

export default ApiHelper;