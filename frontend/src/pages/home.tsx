import OverflowLogo from '../assets/images/logo.png';
import Login from '../components/login';
import LoginAnonymously from '../components/login_anon';

import '../assets/scss/home.scss';

export default function Home() {
    return <div className="Home">
        <img className="Home__Logo" src={OverflowLogo} alt="Overflow logo" />
        <h1>Cards Against Overflow</h1>
        <div className="Home__Login">
            <Login />
            <LoginAnonymously />
        </div>
    </div>;
}