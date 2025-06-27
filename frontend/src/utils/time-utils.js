export const formatDate = (dateString) => {
  if (!dateString) return 'Unknown date';
  const date = new Date(dateString);
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');
  return `${year}-${month}-${day}`;
};

export const formatTime = (dateString) => {
  if (!dateString) return 'Unknown time';
  const date = new Date(dateString);
  return date.toLocaleTimeString('hu-HU', { hour: '2-digit', minute: '2-digit' });
};
