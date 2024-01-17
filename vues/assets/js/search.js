// Open Search
document.querySelector("#search-box").addEventListener("click", function () {
    openSearch();
});

// Close Search
document.querySelectorAll("#search-overlay, #close-search").forEach(el => {
    el.addEventListener("click", function () {
        closeSearch();
    });
});

// Show results
document.querySelector("#input-search").addEventListener("keyup", function () {
    showMatches();
});


function showMatches() {
    let wanted = document.querySelector("#input-search").value;
    // wanted = wanted.replace(/[^a-zA-Z0-9]/g, '');
    let resCount = 0;

    document.querySelectorAll(".search-info").forEach(el => {
        el.innerHTML = el.innerHTML.replaceAll('<strong class="found">', '');
        el.innerHTML = el.innerHTML.replaceAll('</strong>', '');

        let regex = new RegExp(wanted, "gi");
        if (wanted.length > 0 && regex.test(el.innerHTML)) {
            el.innerHTML = el.innerHTML.replace(regex, function (match) {
                return '<strong class="found">' + match + '</strong>';
            });

            el.parentElement.style.display = "flex";
            resCount++;
        } else {
            el.parentElement.style.display = "none";
        }
    });

    if (resCount != 0 || wanted.length == 0) {
        document.querySelector("#no-matches").style.display = "none";
    } else {
        document.querySelector("#no-matches").style.display = "block";
    }
}


document.querySelector("#input-search").addEventListener("input", showMatches);
// Acces to result element
document.querySelectorAll(".search-body").forEach(el => {
    el.addEventListener("click", function () {
        closeSearch();
        let id = this.getAttribute("id").split("_")[1];
        goToCard(id);
        openModal(id);
    });
});

function openSearch() {
    document.querySelector("#search-overlay").style.visibility = "visible";
    document.querySelector("#search-overlay").style.opacity = "1";
    document.querySelector("#search-overlay").style.backdropFilter = "blur(6px)";
    document.querySelector("#search-results").style.right = "0";

    // Disable page scrollbar
    document.body.style.overflow = "hidden";
}

function goToCard(id) {
    document.querySelector("#card_" + id).scrollIntoView({ behavior: 'smooth' });
}

function openModal(id) {
    document.querySelector("#see_" + id).click()
}

function closeSearch() {
    document.querySelector("#search-overlay").style.visibility = "hidden";
    document.querySelector("#search-overlay").style.backdropFilter = "blur(0px)";
    document.querySelector("#search-overlay").style.opacity = "0";

     if (window.innerWidth > 1300) {
        document.querySelector("#search-results").style.right = "-35%";
    } else if (window.innerWidth > 1025) {
        document.querySelector("#search-results").style.right = "-45%";
    }  else if (window.innerWidth > 768) {
        document.querySelector("#search-results").style.right = "-55%";
    } else if (window.innerWidth > 480) {
        document.querySelector("#search-results").style.right = "-85%";
    }else{
        document.querySelector("#search-results").style.right = "-105%";
    }

    // Enable page scrollbar
    document.body.style.overflow = "auto";
}
