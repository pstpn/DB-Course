import csv
import matplotlib.pyplot as plt
from matplotlib.lines import Line2D
import numpy as np
from matplotlib.backends.backend_pdf import PdfPages


class Graph:

    def __init__(self):
        self.sizes = np.array([])
        self.postgres = np.array([])
        self.mongo = np.array([])

    def readFile(self, filename):
        with open(filename) as file:
            spamreader = csv.reader(file, delimiter=';')
            for values in spamreader:
                self.sizes = np.append(self.sizes, int(values[0]))
                self.postgres = np.append(self.postgres, int(values[1]))
                self.mongo = np.append(self.mongo, int(values[2]))

    def buildGraph(self, filename, ylabel):
        with PdfPages(f'./{filename}.pdf') as pdf:
            fig = plt.figure(figsize=(12, 8))
            splt = fig.add_subplot()
            plt.xlabel("Размер изображения (Кбайт)")
            plt.ylabel(ylabel)

            custom_lines = [Line2D([0], [0], color="red", lw=2, marker="x", linestyle="--",  markersize=10),
                            Line2D([0], [0], color="green", lw=2, marker="o", linestyle="solid",  markersize=10)]

            plt.legend(custom_lines, ['PostgreSQL', 'MongoDB'])
            # plt.semilogx()
            plt.semilogy()

            splt.plot(self.sizes, self.postgres, color="red", marker="x", linestyle="--", linewidth=2, markersize=10)
            splt.plot(self.sizes, self.mongo, color="green", marker="o", linestyle="solid", linewidth=2, markersize=10)
            splt.grid(True, linestyle='--', linewidth=0.5, color='gray')
            splt.grid(True)

            pdf.savefig()

            plt.close()


if __name__ == "__main__":

    files = [
        'result1.csv',
        'result2.csv',
    ]

    graph = Graph()
    graph.readFile(files[0])
    graph.buildGraph('measure1', "Время получения изображения (мкс)")

    graph = Graph()
    graph.readFile(files[1])
    graph.buildGraph('measure2', "Время сохранения изображения (мкс)")
