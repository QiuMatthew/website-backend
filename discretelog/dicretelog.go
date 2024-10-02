package discretelog

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strconv"
)

func DiscreteLogHandler(w http.ResponseWriter, r *http.Request) {
	// solve cors error
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	g := r.URL.Query().Get("g")
	h := r.URL.Query().Get("h")
	p := r.URL.Query().Get("p")

	if g == "" || h == "" || p == "" {
		http.Error(w, "Missing Query Parameters", http.StatusBadRequest)
		return
	}

	// check the submitted input type
	if _, err := strconv.ParseInt(g, 10, 64); err != nil {
		http.Error(w, "Inputs must be integers", http.StatusBadRequest)
		return
	}
	if _, err := strconv.ParseInt(h, 10, 64); err != nil {
		http.Error(w, "Inputs must be integers", http.StatusBadRequest)
		return
	}
	if _, err := strconv.ParseInt(p, 10, 64); err != nil {
		http.Error(w, "Inputs must be integers", http.StatusBadRequest)
		return
	}

	// fmt.Println(g, h, p)
	//cmd := exec.Command("python3", "discretelog/script.py", "--generator", g, "--element", h, "--modulus", p)
	cmd := exec.Command("python3", "discretelog/pohlig_hellman.py", "--generator", g, "--element", h, "--modulus", p)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error running script: %v", err)
		http.Error(w, fmt.Sprintf("Error running script %v", err), http.StatusInternalServerError)
		return
	}

	//fmt.Fprintf(w, "Output of script: \n%s", output)
	fmt.Fprintf(w, "%s", output)
}
