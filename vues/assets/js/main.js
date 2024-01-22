document.addEventListener("DOMContentLoaded", function(){
    var modalImage = document.querySelector("#group-image");
    var modalInfos = document.querySelector("#group-infos");
 
    var screenWidth = window.matchMedia("(min-width: 968px)");
    modalImage.onload = function() {
        resizeModalInfos(screenWidth);
    }
 
    screenWidth.addEventListener("change", function(e) {
        resizeModalInfos(e);
    });

    // Initialisation de la carte
    var map = L.map('modal-map').setView([0, 180], 0);
    L.tileLayer('https://tile.openstreetmap.org/{z}/{x}/{y}.png', {
        maxZoom: 19,
        attribution: '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>'
    }).addTo(map);
    var markers = L.layerGroup().addTo(map);

    // Afficher le Modal
    var cards = document.querySelectorAll(".artist-card");
    cards.forEach(card => {
        // Evénement clic du bouton "VOIR"
        card.querySelector(".voir-zone").addEventListener("click", function(){
            // Ajout des données de l'artiste dans le modal
            document.querySelector("#group-image").src = card.querySelector(".artist-image").src;
            document.querySelector("#modal-name").innerHTML = card.querySelector(".artist-name").innerHTML;
            document.querySelector("#modal-creation_date").innerHTML = card.querySelector(".artist-year").innerHTML;
            document.querySelector("#modal-members").innerHTML = card.querySelector(".artist-members").innerHTML;
            document.querySelector("#modal-first_album").innerHTML = card.querySelector(".artist-firstAlbum").innerHTML;
            document.querySelector("#modal-concerts").innerHTML = card.querySelector(".artist-relations").innerHTML;
            // document.querySelector("#modal-map").innerHTML = card.querySelector(".artist-locations").getAttribute("data-location");
            
            // Affichange du modal
            openModal();
            
            // Géolocalisation
            // displayMap(card);
        });
    });
    
    document.addEventListener("click", function(event){
        if (event.target.matches("#modal-close") || event.target.matches("#modal")) {
            closeModal();
        }
    });

    document.querySelectorAll("#see-artists, #arrow-scroll-down").forEach(function(element) {
        element.addEventListener("click", function(){
            document.querySelector("#artists").scrollIntoView({behavior: 'smooth'});
        });
     });
     


    // Fonctions
    // ---------------------------------------------------------------------------
    function resizeModalInfos(e) {
        if (e.matches) {
            modalInfos.style.height = modalImage.offsetHeight+"px";
        }
    }

    function openModal() {
        document.querySelector("#modal").style.visibility = "visible";
        document.querySelector("#modal").style.opacity = "1";
        document.querySelector("#modal").style.backdropFilter = "blur(6px)";
        document.body.style.overflow = "hidden";
    }

    function closeModal() {
        document.querySelector("#modal").style.visibility = "hidden";
        document.querySelector("#modal").style.opacity = "0";
        document.querySelector("#modal").style.backdropFilter = "blur(0px)";
        document.body.style.overflow = "auto";
    }

    
    // ---------------------------------------------------------------------------
    // Geocoding function
    function Geocoding(address) {
        var output = null;
        const token = "pk.2f226e3203c56af77056c10e939e3a0a";
        const settings = {
            async: false,
            crossDomain: true,
            url: 'https://us1.locationiq.com/v1/search?key='+token+'&q='+address+'&format=json&',
            method: 'GET'
        };
        
        $.ajax(settings).done(function (response) {
            output = {
                lat : response[0].lat,
                lon : response[0].lon
            };
        });
        return output;
    }

    async function displayMap(card) {
        var places = card.querySelector(".artist-locations").getAttribute("data-location").replaceAll("[","").replaceAll("]","").split(" ");
            
        var coordinates = [];
        for (let p of places) {
            coordinates[p] = await Geocoding(p);
        }

        // Vider les marqueurs
        markers.clearLayers();

        var allMarkers = [];

        // Ajouter les marqueurs à la carte
        for (var name in coordinates) {
            if (coordinates.hasOwnProperty(name)){
                var c = coordinates[name];
                var marker = L.marker([c.lat, c.lon]).addTo(markers).bindPopup(name);
                allMarkers.push(marker);
            }
        }

        // Créez un objet LatLngBounds vide
        var bounds = L.latLngBounds();

        // Parcourez tous les marqueurs et étendiez les limites à chaque marqueur
        for (var i = 0; i < allMarkers.length; i++) {
            var latlng = allMarkers[i].getLatLng();
            bounds.extend(latlng);
        }

        // Ajustez la vue de la carte pour inclure toutes les limites
        map.fitBounds(bounds);
    }
 
 });
 