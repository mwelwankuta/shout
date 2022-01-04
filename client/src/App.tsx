import React, { FC } from 'react'
import { BrowserRouter, Switch, Route } from 'react-router-dom'
import { ChakraProvider } from '@chakra-ui/react'

import Home from './routes/Home'
import Login from './routes/Login'
import SignUp from './routes/SignUp'

interface Props {

}

const App: FC<Props> = ({ }) => {
    return (
        <ChakraProvider>
            <BrowserRouter>
                <Switch>
                    <Route path="/" component={Home} />
                    <Route path="/login" component={Login} />
                    <Route path="/signup" component={SignUp} />
                </Switch>
            </BrowserRouter>
        </ChakraProvider>
    )
}

export default App
