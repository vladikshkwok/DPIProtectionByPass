function enableEditMode(button) {
    var domainItem = button.closest('.domain-item');
    var domainName = domainItem.querySelector('.domain-name');
    var input = domainItem.querySelector('.domain-edit-input');
    var form = domainItem.querySelector('.edit-form');
    var submitButton = domainItem.querySelector('form button[type="submit"]');

    if (domainName && input && submitButton) {
        input.style.display = 'inline-block';
        input.value = domainName.textContent.trim();  // Copy the current domain name to the input
        domainName.style.display = 'none';
        submitButton.style.display = 'inline-block';
        form.style.display = 'flex'
        input.focus();
        input.setSelectionRange(input.value.length, input.value.length);  // Place cursor at end of input
    } else {
        console.error("Some elements were not found:", domainName, input, submitButton);
    }
}