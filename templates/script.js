var currentStep = 3;

function skipTutorial() {
    window.location.replace("default.html");
}

function submitUserAnswer() {
    var userInput = document.getElementById('userInput').value;
    console.log("Réponse de l'utilisateur :", userInput);

    // Cacher le texte précédent de la section actuelle
    var previousText = document.getElementById('additionalText' + currentStep);
    previousText.style.display = 'none';

    // Afficher le nouveau texte ou effectuer d'autres actions en fonction de la réponse de l'utilisateur
    currentStep++;
    var nextText = document.getElementById('additionalText' + currentStep);
    if (nextText) {
        nextText.style.display = 'block';
    }
}

// Fonction pour gérer la réponse de l'utilisateur
function submitAnswer(answer) {
    console.log("Réponse de l'utilisateur :", answer);

    // Cacher le texte précédent de la section "additionalText"
    
    var additionalText2 = document.getElementById('additionalText2');
    additionalText2.style.display = 'none';

    var additionalText3 = document.getElementById('additionalText3');
    additionalText3.style.display = 'none';

    var additionalText4 = document.getElementById('additionalText4');
    additionalText4.style.display = 'none';

    var additionalText5 = document.getElementById('additionalText5');
    additionalText5.style.display = 'none';

    var additionalText6 = document.getElementById('additionalText6');
    additionalText6.style.display = 'none';

    var additionalText7 = document.getElementById('additionalText7');
    additionalText7.style.display = 'none';

    var resultText = document.getElementById('resultText');
    resultText.style.display = 'none';

    if (answer === 'Five') {
        // Afficher le nouveau texte
        var firstPart = document.getElementById('firstPart');
            if (firstPart) {
                firstPart.style.display = 'none';
            }
        additionalText2.style.display = 'block';
    } else if (answer === 'One' || answer === 'Three') {
        // Logique pour les réponses 'One' et 'Three'
        resultText.textContent = 'Je suis désolé, la réponse était 5';
        resultText.style.display = 'block';
    }
    if (answer === 'Skip') {
        additionalText3.style.display = 'block';
    }
    if (answer === 'a') {
        additionalText4.style.display = 'block';
    }
    if (answer === 'p') {
        additionalText5.style.display = 'block';
    }
    if (answer === 'l') {
        additionalText6.style.display = 'block';
    }
    if (answer === 'e') {
        additionalText7.style.display = 'block';
    }
}
