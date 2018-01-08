function addItem(message) {
    const item = document.createElement('div');
    item.textContent = message;
    item.className = 'item alert';
    document.getElementsByClassName('root')[0].appendChild(item);
}

for (let index = 0; index < 100; index++) {
    addItem('test '+index);
}
