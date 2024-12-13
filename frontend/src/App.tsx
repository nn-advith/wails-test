import {useState} from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import {Greet} from "../wailsjs/go/main/App";
import {LogDebug, LogError, WindowMinimise} from "../wailsjs/runtime/runtime"
import {Encode, Decode} from "../wailsjs/go/codec/Codec"
import { Quit, WindowToggleMaximise } from '../wailsjs/runtime/runtime';

import MinIcon from './assets/icons/window-minimize-solid.svg'
import MaxIcon from './assets/icons/window-maximize-solid.svg'
import CloseIcon from './assets/icons/xmark-solid.svg'

function App() {
    const [resultText, setResultText] = useState("");
    const [text, setText] = useState('');

    const [codecMode, setCodecMode] = useState("encode")

    const updateText = (e: any) => setText(e.target.value);
    
    const updateResultText = (result: string) => setResultText(result);

    function exec() {
        if (codecMode == "encode"){
            Encode(text).then((sEnc) => {
                updateResultText(sEnc)
            }).catch((e) => LogError(e))
        }else {
            Decode(text).then((sEnc) => {
                updateResultText(sEnc)
            }).catch((e) => LogError(e))
        }
  
           
        
    }

    function toggleCodeMode() {
        if (codecMode == "encode"){
            setCodecMode("decode")
        }else {
            setCodecMode("encode")
        }
    }

    

    return (
        <div id="App">
            <div id="control-panel" >
                <div className="minimise" onClick={WindowMinimise}>
                    <img src={MinIcon} width={20} height={20} />
                </div>
                <div className="maximise" onClick={WindowToggleMaximise}>
                    <img src={MaxIcon} width={20} height={20}/>
                </div>
                <div className="quit" onClick={Quit}>
                    <img src={CloseIcon} width={20} height={20} />
                </div>
            </div>
            <div id="content">
                <div className='buttontoggle'>
                    <div className='tgbutton' style={{backgroundColor: codecMode=="encode" ? "white" : "#333", color: codecMode=="encode" ? "#333" : "#fff"}} onClick={() => setCodecMode("encode")}>Encode</div>
                    <div className='tgbutton' style={{backgroundColor: codecMode=="decode" ? "white" : "#333", color: codecMode=="decode" ? "#333" : "#fff"}} onClick={() => setCodecMode("decode")}>Decode</div>
                </div>
                <div className='txtcols'>
                <textarea value={text} rows={10} cols={50} onChange={updateText} className='txtarea'/>
                <textarea value={resultText} rows={10} cols={50} className='txtarea'/>
                </div>
                <div id="input" className="input-box">
                    <button className="btn" onClick={exec}>{codecMode == "encode" ? "Encode" : "Decode"}</button>
                </div>
            </div>
        </div>
    )
}

export default App
