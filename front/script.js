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

const viewAll = document.querySelectorAll('.view-all')
viewAll.forEach((i) => {
    i.addEventListener("mouseover", () => {
        i.style.color = '#0d734f';
    });
    i.addEventListener("mouseout", () => {
        i.style.color = '#12A370'
    });
})

const heads = document.querySelectorAll('.head')
heads.forEach((i) => {
    const originalColor = getComputedStyle(i).color;
    i.addEventListener("mouseover", () => {
        i.style.color = '#12A370';
    });
    i.addEventListener("mouseout", () => {
        i.style.color = originalColor;
    });
})

const prof = document.querySelector('.profile-btn, .add-plan');
prof.addEventListener('mouseover', () => {
    prof.style.backgroundColor = '#ebebebc0';
})
prof.addEventListener('mouseout', () => {
    prof.style.backgroundColor = 'white';
})


const filter = document.getElementById('filter-btn');
const filterSidebar = document.getElementById('filter-sidebar');
const closeBtn = document.getElementById('close-btn');
const overlay = document.getElementById('overlay');

filter.addEventListener('click', () => {
    filterSidebar.classList.add('open');
    overlay.classList.add('active');
});

// Закрытие панели фильтров
closeBtn.addEventListener('click', closeSidebar);
overlay.addEventListener('click', closeSidebar);

function closeSidebar() {
    filterSidebar.classList.remove('open');
    overlay.classList.remove('active');
}

