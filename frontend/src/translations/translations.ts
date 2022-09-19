import i18n from 'i18next';
import { initReactI18next } from 'react-i18next';

const resources = {
    en: {
        translations: {
            login: {
                remote_display: 'Remote display',
                username: 'Username',
                room_id: 'Room ID (Optional)',
                create_join: 'Create / Join',
            },
        },
    },
    fr: {
        translations: {
            login: {
                remote_display: 'Affichage déporté',
                username: 'Pseudo',
                room_id: 'Code salle (Optionel)',
                create_join: 'Créer / Rejoindre',
            }
        },
    },
};

i18n.use(initReactI18next).init({
    resources,
    lng: 'fr',
    interpolation: {
        escapeValue: false,
    },
})

export default i18n;