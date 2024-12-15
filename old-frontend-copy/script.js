//auth-check logika
document.addEventListener('DOMContentLoaded', () => {
    const currentPage = window.location.pathname.split('/').pop();
    const isLoggedIn = localStorage.getItem('isLoggedIn');

    if (currentPage === 'index.html' && !isLoggedIn) {
        window.location.href = 'auth.html';
    }

    if (currentPage === 'auth.html' && isLoggedIn) {
        window.location.href = 'index.html';
    }
});

//auth logika
document.addEventListener('DOMContentLoaded', () => {
    const loginForm = document.querySelector('form');
    if (loginForm) {
        loginForm.addEventListener('submit', (event) => {
            event.preventDefault();

            const email = document.querySelector('#email').value; 
            const password = document.querySelector('#password').value; 

            if (email === 'test@mail.ru' && password === '123') {
                localStorage.setItem('isLoggedIn', true);
                window.location.href = 'index.html';
            } else {
                alert('Неверный email или пароль');
            }
        });
    }
});

//outlog logika
document.querySelector('.logout-btn')?.addEventListener('click', () => {
    localStorage.removeItem('isLoggedIn');
    window.location.href = 'auth.html';
});

//krut' (btn-animation)
const btnSmall = document.querySelectorAll('.btn-small, .btn, .btn-outline');
btnSmall.forEach((btn) => {
    btn.addEventListener("mouseover", () => {
        btn.style.transform = 'scale(1.1)';
        btn.style.transition = 'transform 0.3s ease';
        btn.style.boxShadow = '2px 5px 7px rgba(0, 0, 0, 0.1)';
    });
    btn.addEventListener("mouseout", () => {
        btn.style.transform = 'scale(1)';
        btn.style.transition = 'transform 0.3s ease';
        btn.style.boxShadow = '0 0 0';
    });
});

//ne krut' (link-color-swth)
const viewAll = document.querySelectorAll('.view-all');
viewAll.forEach((i) => {
    i.addEventListener("mouseover", () => {
        i.style.color = '#0d734f';
    });
    i.addEventListener("mouseout", () => {
        i.style.color = '#12A370';
    });
});

//krut' (heads-animation)
const heads = document.querySelectorAll('.head');
heads.forEach((i) => {
    const originalColor = getComputedStyle(i).color;
    i.addEventListener("mouseover", () => {
        i.style.color = '#12A370';
    });
    i.addEventListener("mouseout", () => {
        i.style.color = originalColor;
    });
});

//krut' (profile-btn-animation)
const prof = document.querySelector('.profile-btn, .add-plan');
prof?.addEventListener('mouseover', () => {
    prof.style.backgroundColor = '#ebebebc0';
});
prof?.addEventListener('mouseout', () => {
    prof.style.backgroundColor = 'white';
});

//dayte mne chertovi filtres (filters-open-logika)
const filter = document.getElementById('filter-btn');
const filterSidebar = document.getElementById('filter-sidebar');
const closeBtn = document.getElementById('close-btn');
const overlay = document.getElementById('overlay');

filter?.addEventListener('click', () => {
    filterSidebar.classList.add('open');
    overlay.classList.add('active');
});

//nahuy eti filtri (filters-close-logika)
closeBtn?.addEventListener('click', closeSidebar);
overlay?.addEventListener('click', closeSidebar);

function closeSidebar() {
    filterSidebar.classList.remove('open');
    overlay.classList.remove('active');
}