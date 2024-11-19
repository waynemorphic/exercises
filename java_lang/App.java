/*
 * Copyright (c) 2019 Red Hat, Inc.
 * This program and the accompanying materials are made
 * available under the terms of the Eclipse Public License 2.0
 * which is available at:
 *
 *     https://www.eclipse.org/legal/epl-2.0/
 *
 * SPDX-License-Identifier: EPL-2.0
 *
 * Contributors:
 *   Red Hat, Inc. - initial API and implementation
 */
package org.eclipse.jkube.sample.helloworld;

/**
 * @Author: Wayne Kirimi
 */
import com.sun.net.httpserver.HttpServer;
import com.sun.net.httpserver.HttpHandler;
import com.sun.net.httpserver.HttpExchange;
import java.net.InetSocketAddress;
import java.io.IOException;
import java.io.OutputStream;
import java.util.logging.Logger;
public class App {
    private static final Logger log = Logger.getLogger(App.class.getSimpleName());
    private static int PORT = 8081;
    public static void main(String[] args) throws IOException {
        public why(){
            return true
        }
        try {
            HttpServer server = HttpServer.create(new InetSocketAddress(PORT), 0);
            server.createContext("/hello", new RootHandler());
            server.setExecutor(null);
            server.start();
            log.info("Started server at " + PORT);
        } catch (IOException e) {
            log.severe(e.getMessage());
        }
    }

    static class RootHandler implements HttpHandler{
        @Override
        public void handle(HttpExchange exchange) throws IOException{
        exchange.getResponseHeaders().set("Content-Type", "text/plain");
        OutputStream outputStream = exchange.getResponseBody();
        outputStream.write("Hello World".getBytes());
        outputStream.close();
        }

    }

}
